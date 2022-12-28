package k8s

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/limeschool/gin"
	apiApps "k8s.io/api/apps/v1"
	apiCore "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"service/consts"
	"service/errors"
	"service/tools/remote/model"
	"sigs.k8s.io/yaml"
	"time"
)

func (c *Client) CreateDeployment(ctx *gin.Context, config model.ServiceConfig) error {
	j2, err := yaml.YAMLToJSON([]byte(config.Yaml))
	if err != nil {
		return errors.New("yaml格式错误")
	}
	deploy := apiApps.Deployment{}
	if err = json.Unmarshal(j2, &deploy); err != nil {
		return err
	}

	// 重写meta 方便后续创建service
	deploy.ObjectMeta = c.objectMeta(config.Namespace, config.ServiceName)
	deploy.Spec.Selector.MatchLabels = c.selector(config.ServiceName)
	deploy.Spec.Template.ObjectMeta = c.objectMeta(config.Namespace, config.ServiceName)

	// 创建deployment
	client := c.client.AppsV1().Deployments(config.Namespace)
	if _, err = client.Get(ctx, config.ServiceName, meta.GetOptions{}); err != nil {
		_, err = client.Create(ctx, &deploy, meta.CreateOptions{})
	} else {
		_, err = client.Update(ctx, &deploy, meta.UpdateOptions{})
	}
	return err
}

func (c *Client) CreateDaemonSet(ctx *gin.Context, config model.ServiceConfig) error {
	j2, err := yaml.YAMLToJSON([]byte(config.Yaml))
	if err != nil {
		return errors.New("yaml格式错误")
	}
	daemonSet := apiApps.DaemonSet{}
	if err = json.Unmarshal(j2, &daemonSet); err != nil {
		return err
	}

	// 重写meta 方便后续创建service
	daemonSet.ObjectMeta = c.objectMeta(config.Namespace, config.ServiceName)
	daemonSet.Spec.Selector.MatchLabels = c.selector(config.ServiceName)
	daemonSet.Spec.Template.ObjectMeta = c.objectMeta(config.Namespace, config.ServiceName)

	// 创建deployment
	client := c.client.AppsV1().DaemonSets(config.Namespace)
	if _, err = client.Get(ctx, config.ServiceName, meta.GetOptions{}); err != nil {
		_, err = client.Create(ctx, &daemonSet, meta.CreateOptions{})
	} else {
		_, err = client.Update(ctx, &daemonSet, meta.UpdateOptions{})
	}
	return err
}

func (c *Client) DeleteDeployment(ctx *gin.Context, config model.ServiceConfig) error {
	client := c.client.AppsV1().Deployments(config.Namespace)
	return client.Delete(ctx, config.ServiceName, meta.DeleteOptions{})
}

func (c *Client) DeleteDaemonSet(ctx *gin.Context, config model.ServiceConfig) error {
	client := c.client.AppsV1().DaemonSets(config.Namespace)
	return client.Delete(ctx, config.ServiceName, meta.DeleteOptions{})
}

func (c *Client) GetPods(ctx *gin.Context, namespace, service string) ([]apiCore.Pod, error) {
	client := c.client.CoreV1().Pods(namespace)
	list, err := client.List(ctx, meta.ListOptions{LabelSelector: fmt.Sprintf("%v=%v", consts.K8sLabelPrefix, service)})
	if err != nil {
		return nil, err
	}
	return list.Items, err
}

func (c *Client) GetPodsStatus(ctx *gin.Context, namespace, service string) error {
	pods, err := c.GetPods(ctx, namespace, service)
	if err != nil {
		return err
	}
	for _, item := range pods {
		if item.Status.Phase == "Running" {
			continue
		}
		if len(item.Status.ContainerStatuses) <= 0 {
			return errors.New("not have container")
		}
		container := item.Status.ContainerStatuses[0]
		if container.Ready {
			continue
		}
		if container.State.Waiting != nil && container.State.Waiting.Message != "" {
			return errors.NewF("【%v】%v", container.State.Waiting.Reason, container.State.Waiting.Message)
		}
	}
	return nil
}

func (c *Client) GetDaemonSetStatus(ctx *gin.Context, config model.ServiceConfig) (string, error) {
	client := c.client.AppsV1().DaemonSets(config.Namespace)
	deploy, err := client.Get(ctx, config.ServiceName, meta.GetOptions{})
	if err != nil {
		return model.UnAvailable, err
	}

	condList := deploy.Status.Conditions
	if len(condList) != 2 {
		return model.UnAvailable, errors.New("not init condList")
	}

	available := condList[0]
	progressing := condList[1]

	if available.Status == "True" && progressing.Reason == "NewReplicaSetAvailable" {
		return model.Available, nil
	}

	if available.Status == "True" && progressing.Reason == "ReplicaSetUpdated" {
		// 获取pod 判断pod的状态
		if err = c.GetPodsStatus(ctx, config.Namespace, config.ServiceName); err != nil {
			return model.RollUpdateFail, err
		}
		return model.RollUpdate, nil
	}

	// 滚动更新失败
	if available.Status == "True" && progressing.Reason == "ProgressDeadlineExceeded" {
		return model.RollUpdateFail, errors.New(progressing.Message)
	}

	// 初始化
	if available.Status == "False" && progressing.Reason == "MinimumReplicasUnavailable" {
		return model.Update, nil
	}
	// 创建中
	if available.Status == "False" && progressing.Reason == "ReplicaSetUpdated" {
		if err = c.GetPodsStatus(ctx, config.Namespace, config.ServiceName); err != nil {
			return model.UnAvailable, err
		}
		return model.Update, nil
	}
	// 创建失败
	if available.Status == "False" && progressing.Reason == "ProgressDeadlineExceeded" {
		return model.UnAvailable, errors.New(progressing.Message)
	}

	return model.UnAvailable, errors.NewF("未识别的服务状态:%v:%v", available.Status, progressing.Reason)
}

func (c *Client) GetDeploymentStatus(ctx *gin.Context, config model.ServiceConfig) (string, error) {

	client := c.client.AppsV1().Deployments(config.Namespace)
	deploy, err := client.Get(ctx, config.ServiceName, meta.GetOptions{})
	if err != nil {
		return model.UnAvailable, err
	}

	condList := deploy.Status.Conditions
	if len(condList) != 2 {
		return model.UnAvailable, errors.New("not init condList")
	}

	available := condList[0]
	progressing := condList[1]

	if available.Status == "True" && progressing.Reason == "NewReplicaSetAvailable" {
		return model.Available, nil
	}

	if available.Status == "True" && progressing.Reason == "ReplicaSetUpdated" {
		// 获取pod 判断pod的状态
		if err = c.GetPodsStatus(ctx, config.Namespace, config.ServiceName); err != nil {
			return model.RollUpdateFail, err
		}
		return model.RollUpdate, nil
	}

	// 滚动更新失败
	if available.Status == "True" && progressing.Reason == "ProgressDeadlineExceeded" {
		return model.RollUpdateFail, errors.New(progressing.Message)
	}

	// 初始化
	if available.Status == "False" && progressing.Reason == "MinimumReplicasUnavailable" {
		return model.Update, nil
	}
	// 创建中
	if available.Status == "False" && progressing.Reason == "ReplicaSetUpdated" {
		if err = c.GetPodsStatus(ctx, config.Namespace, config.ServiceName); err != nil {
			return model.UnAvailable, err
		}
		return model.Update, nil
	}
	// 创建失败
	if available.Status == "False" && progressing.Reason == "ProgressDeadlineExceeded" {
		return model.UnAvailable, errors.New(progressing.Message)
	}

	return model.UnAvailable, errors.NewF("未识别的服务状态:%v:%v", available.Status, progressing.Reason)
}

func (c *Client) DeleteService(ctx *gin.Context, srv model.ServiceConfig) error {
	if srv.Type == consts.K8sControllerDeployment {
		return c.DeleteDeployment(ctx, srv)
	}
	if srv.Type == consts.K8sControllerDeployment {
		return c.DeleteDaemonSet(ctx, srv)
	}
	return errors.New("错误的服务控制器类型")
}

func (c *Client) CreateService(ctx *gin.Context, srv model.ServiceConfig) error {
	if srv.Type == consts.K8sControllerDeployment {
		return c.CreateDeployment(ctx, srv)
	}
	if srv.Type == consts.K8sControllerDeployment {
		return c.CreateDaemonSet(ctx, srv)
	}
	return errors.New("错误的服务控制器类型")
}

func (c *Client) GetServiceRelease(ctx *gin.Context, srv model.ServiceConfig) error {
	for {
		switch srv.Type {
		case consts.K8sControllerDeployment:
			if status, err := c.GetDeploymentStatus(ctx, srv); status != model.Update && status != model.RollUpdate {
				return err
			}
		case consts.K8sControllerDaemonSet:
			if status, err := c.GetDaemonSetStatus(ctx, srv); status != model.Update && status != model.RollUpdate {
				return err
			}
		default:
			return errors.New("不支持的服务运行类型")
		}

		time.Sleep(1 * time.Second)
	}
}

func (c *Client) GetServicePods(ctx *gin.Context, config model.ServiceConfig) ([]model.Pod, error) {
	list, err := c.GetPods(ctx, config.Namespace, config.ServiceName)
	if err != nil {
		return nil, err
	}

	var respList []model.Pod
	for _, item := range list {
		if len(item.Spec.Containers) <= 0 || len(item.Status.ContainerStatuses) <= 0 {
			continue
		}

		status := item.Status.ContainerStatuses[0]
		errStr := ""
		if status.State.Waiting != nil && status.State.Waiting.Message != "" {
			errStr = fmt.Sprintf("【%v】%v", status.State.Waiting.Reason, status.State.Waiting.Message)

		}

		temp := model.Pod{
			Name:        item.Name,
			Status:      string(item.Status.Phase),
			Image:       item.Spec.Containers[0].Image,
			Err:         errStr,
			Node:        item.Spec.NodeName,
			CreatedTime: item.Status.StartTime.String(),
			Restart:     int(status.RestartCount),
		}
		respList = append(respList, temp)
	}
	return respList, nil
}
