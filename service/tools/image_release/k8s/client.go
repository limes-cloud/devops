package k8s

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/limeschool/gin"
	"io"
	apiCore "k8s.io/api/core/v1"
	apiNetworking "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	syaml "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"net"
	"net/http"
	"service/consts"
	"service/errors"
	"service/tools"
	"service/tools/image_release"
	sigyaml "sigs.k8s.io/yaml"
	"time"
)

type Client struct {
	KubernetesClient *kubernetes.Clientset
	DynamicClient    dynamic.Interface
	Namespace        string
	ResourceInfo     schema.GroupVersionKind
}

func NewK8sClient(host, token, namespace string) (*Client, error) {
	var tlsConfig = &tls.Config{
		InsecureSkipVerify: true, // 忽略证书验证
	}
	var transport http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}
	config := &rest.Config{
		Host:        host,
		BearerToken: token,
		Transport:   transport,
	}

	k8sCli, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	dycCli, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		KubernetesClient: k8sCli,
		DynamicClient:    dycCli,
		Namespace:        namespace,
	}, nil
}

func (c *Client) GtGVR(gvk schema.GroupVersionKind) (schema.GroupVersionResource, error) {
	gr, err := restmapper.GetAPIGroupResources(c.KubernetesClient.Discovery())
	if err != nil {
		return schema.GroupVersionResource{}, err
	}

	mapper := restmapper.NewDiscoveryRESTMapper(gr)

	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return schema.GroupVersionResource{}, err
	}

	return mapping.Resource, nil
}

// DeleteFormYaml 删除资源清单
func (c *Client) DeleteFormYaml(ctx *gin.Context, applyYaml string) (err error) {
	d := yaml.NewYAMLOrJSONDecoder(bytes.NewBufferString(applyYaml), 4096)
	for {
		var rawObj runtime.RawExtension
		err = d.Decode(&rawObj)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("decode yaml is err %v", err)
		}

		obj, _, err := syaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		if err != nil {
			return fmt.Errorf("rawobj is err %v", err)
		}

		mapper, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return fmt.Errorf("tounstructured is err %v", err)
		}

		object := &unstructured.Unstructured{Object: mapper}
		c.ResourceInfo = object.GroupVersionKind()
		gvr, err := c.GtGVR(c.ResourceInfo)
		if err != nil {
			return err
		}

		if err != nil {
			return fmt.Errorf("unable to marshal resource as yaml: %w", err)
		}

		// 当为设置namespace 以yaml中的为准
		if c.Namespace == "" {
			c.Namespace = object.GetNamespace()
		}

		namespaceRes := c.DynamicClient.Resource(gvr).Namespace(c.Namespace)
		// 查询不到资源
		if _, err := namespaceRes.Get(ctx, object.GetName(), v1.GetOptions{}); err != nil {
			return nil
		}

		if err = namespaceRes.Delete(ctx, object.GetName(), v1.DeleteOptions{}); err != nil {
			return fmt.Errorf("delete to patch resource: %w", err)
		}

	}
	return nil
}

// UpdateFromYaml 执行yaml
func (c *Client) UpdateFromYaml(ctx *gin.Context, applyYaml string) (err error) {
	d := yaml.NewYAMLOrJSONDecoder(bytes.NewBufferString(applyYaml), 4096)
	for {
		var rawObj runtime.RawExtension
		err = d.Decode(&rawObj)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("decode yaml is err %v", err)
		}

		obj, _, err := syaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		if err != nil {
			return fmt.Errorf("rawobj is err %v", err)
		}

		mapper, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return fmt.Errorf("tounstructured is err %v", err)
		}

		object := &unstructured.Unstructured{Object: mapper}
		c.ResourceInfo = object.GroupVersionKind()
		gvr, err := c.GtGVR(c.ResourceInfo)
		if err != nil {
			return err
		}
		unstructuredYaml, err := sigyaml.Marshal(object)
		if err != nil {
			return fmt.Errorf("unable to marshal resource as yaml: %w", err)
		}

		// 当为设置namespace 以yaml中的为准
		if c.Namespace == "" {
			c.Namespace = object.GetNamespace()
		}

		namespaceRes := c.DynamicClient.Resource(gvr).Namespace(c.Namespace)
		// 获取不到资源则进行创建资源
		info, getErr := namespaceRes.Get(ctx, object.GetName(), v1.GetOptions{})
		if getErr != nil {
			_, err = namespaceRes.Create(ctx, object, v1.CreateOptions{})
			return err
		}

		// 获取状态
		status := c.GetStatus(info.Object)
		if !status.Available && status.Progressing {
			return errors.New("资源正在创建中")
		}

		if !status.Available {
			if err := c.DeleteFormYaml(ctx, applyYaml); err != nil {
				return errors.NewF("清理异常资源失败：%v", err)
			}
			_, err = namespaceRes.Create(ctx, object, v1.CreateOptions{})
			return err
		}

		// 获取到资源则进行修改
		force := true
		if _, err := namespaceRes.Patch(ctx, object.GetName(), types.ApplyPatchType, unstructuredYaml, v1.PatchOptions{
			FieldManager: object.GetName(),
			Force:        &force,
		}); err != nil {
			return fmt.Errorf("unable to patch resource: %w", err)
		}
	}
	return nil
}

type sourceStatus struct {
	Progressing bool
	Available   bool
}

func (c *Client) GetStatus(info map[string]interface{}) *sourceStatus {
	statusInfo, ok := info["status"].(map[string]any)
	if !ok {
		return nil
	}

	resp := &sourceStatus{}
	list, _ := statusInfo["conditions"].([]any)
	for _, item := range list {
		cond, _ := item.(map[string]any)

		// 是否在线
		if cond["type"] == "Available" {
			resp.Available = cond["status"] != "False"
		}

		// 是否在处理中
		if cond["type"] == "Progressing" {
			resp.Progressing = cond["status"] != "False"
		}
	}
	return resp
}

func (c *Client) GetStartStatus(ctx *gin.Context, srv string) error {
	gvr, err := c.GtGVR(c.ResourceInfo)
	if err != nil {
		return err
	}
	namespaceRes := c.DynamicClient.Resource(gvr).Namespace(c.Namespace)
	for {
		time.Sleep(5 * time.Second)
		// 获取资源
		info, getErr := namespaceRes.Get(ctx, srv, v1.GetOptions{})
		if getErr != nil {
			return err
		}
		status := c.GetStatus(info.Object)

		// 就绪
		if status.Available {
			return nil
		}

		if status.Progressing {
			continue
		}

		// 否则获取pod错误信息
		label := fmt.Sprintf("%v=%v", consts.K8sLabelPrefix, srv)
		list, err := c.KubernetesClient.CoreV1().Pods(c.Namespace).List(ctx, v1.ListOptions{LabelSelector: label})
		if err != nil {
			return err
		}

		if len(list.Items) == 0 && status.Progressing {
			continue
		}

		if len(list.Items) == 0 {
			return errors.New("获取pod信息失败")
		}

		pod := list.Items[0]
		// 读取pod信息
		for _, item := range pod.Status.ContainerStatuses {
			if item.Ready {
				continue
			}
			if item.State.Waiting != nil {
				return errors.NewF("【%v】%v", item.State.Waiting.Reason, item.State.Waiting.Message)
			}
		}
		break
	}
	return nil
}

func (c *Client) AddNetworkSecret(ctx *gin.Context, config image_release.NetworkConfig) error {
	if config.Key == "" || config.Cert == "" {
		return errors.New("key and cert must not empty str")
	}

	var tp apiCore.SecretType = "kubernetes.io/tls"
	secret := c.KubernetesClient.CoreV1().Secrets(config.Namespace)
	_, err := secret.Create(ctx, &apiCore.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      config.Service,
			Namespace: config.Namespace,
			Labels: map[string]string{
				"from": "devops",
				"app":  config.Service,
			},
		},
		Data: map[string][]byte{
			"tls.crt": []byte(config.Cert),
			"tls.key": []byte(config.Key),
		},
		Type: tp,
	}, v1.CreateOptions{})
	return err
}

func (c *Client) AddNetworkService(ctx *gin.Context, config image_release.NetworkConfig) error {
	service := c.KubernetesClient.CoreV1().Services(config.Namespace)
	_, err := service.Create(ctx, &apiCore.Service{
		ObjectMeta: v1.ObjectMeta{
			Name:      config.Service,
			Namespace: config.Namespace,
			Labels: map[string]string{
				"from": "devops",
				"app":  config.Service,
			},
		},
		Status: apiCore.ServiceStatus{},
		Spec: apiCore.ServiceSpec{
			Selector: map[string]string{
				"app": config.Service,
			},
			Ports: []apiCore.ServicePort{
				{
					Port:     int32(config.RunPort),
					Protocol: apiCore.ProtocolTCP,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(config.TargetPort),
					},
				},
			},
			Type: apiCore.ServiceTypeClusterIP,
		},
	}, v1.CreateOptions{})
	return err
}

func (c *Client) AddNetworkIngress(ctx *gin.Context, config image_release.NetworkConfig, isTls bool) error {
	var tls []apiNetworking.IngressTLS
	if isTls {
		tls = append(tls, apiNetworking.IngressTLS{
			Hosts:      []string{config.Host},
			SecretName: config.Service,
		})
	}
	var pathType apiNetworking.PathType = "Prefix"
	ingress := c.KubernetesClient.NetworkingV1().Ingresses(config.Namespace)

	_, err := ingress.Create(ctx, &apiNetworking.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Namespace: config.Namespace,
			Name:      config.Service,
			Annotations: map[string]string{
				"ingress.kubernetes.io/ssl-redirect": fmt.Sprint(config.Redirect),
			},
			Labels: map[string]string{
				"from": "devops",
				"app":  config.Service,
			},
		},
		Spec: apiNetworking.IngressSpec{
			IngressClassName: tools.String("nginx"),
			TLS:              tls,
			Rules: []apiNetworking.IngressRule{
				{
					Host: config.Host,
					IngressRuleValue: apiNetworking.IngressRuleValue{
						HTTP: &apiNetworking.HTTPIngressRuleValue{
							Paths: []apiNetworking.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: apiNetworking.IngressBackend{
										Service: &apiNetworking.IngressServiceBackend{
											Name: config.Service,
											Port: apiNetworking.ServiceBackendPort{
												Number: int32(config.RunPort),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, v1.CreateOptions{})
	return err
}

func (c *Client) DeleteNetworkSecret(ctx *gin.Context, config image_release.NetworkConfig) error {
	secret := c.KubernetesClient.CoreV1().Secrets(config.Namespace)
	return secret.Delete(ctx, config.Service, v1.DeleteOptions{})
}

func (c *Client) DeleteNetworkService(ctx *gin.Context, config image_release.NetworkConfig) error {
	service := c.KubernetesClient.CoreV1().Services(config.Namespace)
	return service.Delete(ctx, config.Service, v1.DeleteOptions{})
}
func (c *Client) DeleteNetworkIngress(ctx *gin.Context, config image_release.NetworkConfig) error {
	ingress := c.KubernetesClient.NetworkingV1().Ingresses(config.Namespace)
	return ingress.Delete(ctx, config.Service, v1.DeleteOptions{})
}

func (c *Client) AddNetwork(ctx *gin.Context, config image_release.NetworkConfig) error {
	//生成yaml
	isTls := false

	// 创建service
	if err := c.AddNetworkService(ctx, config); err != nil {
		return err
	}

	// 判断是否需要生成tls,提前生成secret
	if config.Key != "" && config.Cert != "" {
		if err := c.AddNetworkSecret(ctx, config); err != nil {
			return err
		}
		isTls = true
	}

	if err := c.AddNetworkIngress(ctx, config, isTls); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteNetwork(ctx *gin.Context, config image_release.NetworkConfig) error {
	_ = c.DeleteNetworkSecret(ctx, config)
	_ = c.DeleteNetworkService(ctx, config)
	_ = c.DeleteNetworkIngress(ctx, config)
	return nil
}
