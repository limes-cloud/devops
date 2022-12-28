package service

import (
	"dc/errors"
	"dc/tools"
	"dc/tools/exec"
	"dc/types"
	"encoding/json"
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/limeschool/gin"
	"os"
	"sigs.k8s.io/yaml"
)

type Container struct {
	Image   string `json:"image"`
	Command string `json:"command"`
	Created int64  `json:"created"`
	State   string `json:"state"`
	Status  string `json:"status"`
}

func GetServiceRelease(ctx *gin.Context, in *types.GetServiceReleaseRequest) (map[string]any, error) {
	pods, err := GetServicePods(ctx, &types.GetServicePodsRequest{
		ServiceName: in.ServiceName,
		Replicas:    in.Replicas,
	})

	if err != nil {
		return nil, err
	}
	status := true
	state := "running"
	for _, item := range pods {
		if item.State != "running" {
			status = false
			state = item.State
		}
	}

	return map[string]any{
		"status": status,
		"state":  state,
	}, nil
}

func GetServicePods(ctx *gin.Context, in *types.GetServicePodsRequest) ([]Container, error) {
	var resp []Container
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	var names []string

	if in.Replicas == 1 {
		names = append(names, in.ServiceName)
	} else {
		for i := 0; i < in.Replicas; i++ {
			names = append(names, fmt.Sprintf("%v-%v", in.ServiceName, i))
		}
	}

	config := docker.ListContainersOptions{All: true, Size: true, Filters: map[string][]string{"name": names}}
	containers, err := client.ListContainers(config)
	if err != nil {
		return nil, err
	}
	for _, item := range containers {
		resp = append(resp, Container{
			Image:   item.Image,
			Command: item.Command,
			Created: item.Created,
			State:   item.State,
			Status:  item.Status,
		})
	}
	return resp, nil
}

type DockerCompose struct {
	Version  string                    `json:"version"`
	Services map[string]map[string]any `json:"services"`
}

func AddService(ctx *gin.Context, in *types.AddServiceRequest) error {
	j2, err := yaml.YAMLToJSON([]byte(in.Yaml))
	if err != nil {
		return errors.New("yaml格式错误")
	}

	var m DockerCompose
	if err = json.Unmarshal(j2, &m); err != nil {
		return errors.New("yaml格式错误")
	}

	if len(m.Services) != 1 && in.Replicas != 1 {
		return errors.New("副本数不为1时不支持部署多服务，请检查docker-compose文件")
	}

	// 重写部分值
	if in.Replicas != 1 {
		srv := ""
		info := map[string]any{}
		// 获取服务的key 信息
		for key, item := range m.Services {
			srv = key
			info = item
		}

		delete(m.Services, srv)

		// 添加副本
		for i := 0; i < in.Replicas; i++ {
			temp := DeepCopyByJson(info)
			// 重写值 container_name service_name port
			name := fmt.Sprintf("%v-%v", srv, i)
			temp["container_name"] = name
			temp["ports"] = []string{fmt.Sprintf("%v:%v", fmt.Sprint(in.RunPort+i), fmt.Sprint(in.ListenPort))}
			m.Services[name] = temp
		}
	}

	// 重新生成yaml
	byteData, _ := json.Marshal(m)
	yamlByte, _ := yaml.JSONToYAML(byteData)

	// 保存文件
	workDir := ctx.Config.GetString("work_dir") + "/" + in.ServiceName
	if !tools.FileExist(workDir) && os.MkdirAll(workDir, os.ModePerm) != nil {
		return errors.New("创建服务工作目录失败")
	}
	if os.WriteFile(workDir+"/docker-compose.yaml", yamlByte, os.ModePerm) != nil {
		return errors.New("docker-compose文件失败")
	}

	// 判断环境是否正常
	e := exec.New()
	execType := ctx.Config.GetDefaultString("exec_type", "/bin/sh")
	if err = CheckEnv(e, execType); err != nil {
		return err
	}

	// 登陆docker
	loginShell := fmt.Sprintf("docker login -u %v -p %v %v", in.ImageUser, in.ImagePass, in.ImageRegistry)
	cmd := e.Command(execType, "-c", loginShell)
	if _, err = cmd.CombinedOutput(); err != nil {
		return errors.New("docker-compose 环境检测异常")
	}

	// 执行服务上线
	cmd = e.Command(execType, "-c", "docker-compose up -d")
	cmd.SetDir(workDir)
	if resByte, err := cmd.CombinedOutput(); err != nil {
		return errors.New(string(resByte))
	}

	return nil

}

func CheckEnv(e exec.Interface, execType string) error {

	// 检测docker
	cmd := e.Command(execType, "-c", "docker -v")
	if _, err := cmd.CombinedOutput(); err != nil {
		return errors.New("docker 环境检测异常")
	}

	// 检测docker-compose
	cmd = e.Command(execType, "-c", "docker-compose -v")
	if _, err := cmd.CombinedOutput(); err != nil {
		return errors.New("docker-compose 环境检测异常")
	}

	return nil
}

func DeleteService(ctx *gin.Context, in *types.DeleteServiceRequest) error {
	e := exec.New()
	execType := ctx.Config.GetDefaultString("exec_type", "/bin/sh")
	if err := CheckEnv(e, execType); err != nil {
		return err
	}

	// 进行服务下线
	workDir := ctx.GetString("work_dir") + "/" + in.ServiceName
	cmd := e.Command(execType, "-c", "docker-compose down")
	cmd.SetDir(workDir)
	if resByte, err := cmd.CombinedOutput(); err != nil {
		return errors.New(string(resByte))
	}
	return nil
}

func DeepCopyByJson(src map[string]any) map[string]any {
	var dst = make(map[string]any)
	b, _ := json.Marshal(src)
	_ = json.Unmarshal(b, &dst)
	return dst
}
