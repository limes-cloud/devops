package k8s

import (
	"fmt"
	"github.com/limeschool/gin"
	apiCore "k8s.io/api/core/v1"
	apiNetworking "k8s.io/api/networking/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"service/errors"
	"service/tools"
	"service/tools/remote/model"
)

// ingressRules 获取ingress
func (c *Client) ingressRules(config model.NetworkConfig) []apiNetworking.IngressRule {
	var pathType apiNetworking.PathType = "Prefix"
	return []apiNetworking.IngressRule{
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
									Name: config.ServiceName,
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
	}
}

// CreateNetworkSecret 添加网路配置的密钥
func (c *Client) CreateNetworkSecret(ctx *gin.Context, config model.NetworkConfig) error {
	if config.Key == "" || config.Cert == "" {
		return errors.New("key and cert must not empty str")
	}

	var tp apiCore.SecretType = "kubernetes.io/tls"
	secret := c.client.CoreV1().Secrets(config.Namespace)
	_, err := secret.Create(ctx, &apiCore.Secret{
		ObjectMeta: c.objectMeta(config.Namespace, config.ServiceName),
		Data: map[string][]byte{
			"tls.crt": []byte(config.Cert),
			"tls.key": []byte(config.Key),
		},
		Type: tp,
	}, meta.CreateOptions{})
	return err
}

// CreateNetworkService 创建服务的service
func (c *Client) CreateNetworkService(ctx *gin.Context, config model.NetworkConfig) error {
	service := c.client.CoreV1().Services(config.Namespace)
	srvConf := apiCore.Service{
		ObjectMeta: c.objectMeta(config.Namespace, config.ServiceName),
		Status:     apiCore.ServiceStatus{},
		Spec: apiCore.ServiceSpec{
			Selector: c.selector(config.ServiceName),
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
	}
	_, err := service.Create(ctx, &srvConf, meta.CreateOptions{})
	return err
}

// CreateNetworkIngress 创建服务的ingress
func (c *Client) CreateNetworkIngress(ctx *gin.Context, config model.NetworkConfig, isTls bool) error {
	var tls []apiNetworking.IngressTLS

	if isTls {
		tls = append(tls, apiNetworking.IngressTLS{
			Hosts:      []string{config.Host},
			SecretName: config.ServiceName,
		})
	}

	ingress := c.client.NetworkingV1().Ingresses(config.Namespace)
	objectMeta := c.objectMeta(config.Namespace, config.ServiceName)
	objectMeta.Annotations = map[string]string{
		"ingress.kubernetes.io/ssl-redirect": fmt.Sprint(config.Redirect),
	}

	ingressConf := apiNetworking.Ingress{
		ObjectMeta: objectMeta,
		Spec: apiNetworking.IngressSpec{
			IngressClassName: tools.String("nginx"),
			TLS:              tls,
			Rules:            c.ingressRules(config),
		},
	}
	_, err := ingress.Create(ctx, &ingressConf, meta.CreateOptions{})
	return err
}

// DeleteNetworkSecret 删除服务的密钥
func (c *Client) DeleteNetworkSecret(ctx *gin.Context, config model.NetworkConfig) error {
	secret := c.client.CoreV1().Secrets(config.Namespace)
	return secret.Delete(ctx, config.ServiceName, meta.DeleteOptions{})
}

// DeleteNetworkService 删除服务的service
func (c *Client) DeleteNetworkService(ctx *gin.Context, config model.NetworkConfig) error {
	service := c.client.CoreV1().Services(config.Namespace)
	return service.Delete(ctx, config.ServiceName, meta.DeleteOptions{})
}

// DeleteNetworkIngress 删除服务的ingress
func (c *Client) DeleteNetworkIngress(ctx *gin.Context, config model.NetworkConfig) error {
	ingress := c.client.NetworkingV1().Ingresses(config.Namespace)
	return ingress.Delete(ctx, config.ServiceName, meta.DeleteOptions{})
}

// CreateNetwork 添加网络配置
func (c *Client) CreateNetwork(ctx *gin.Context, config model.NetworkConfig) error {
	//生成yaml
	isTls := false

	// 创建service
	if err := c.CreateNetworkService(ctx, config); err != nil {
		return err
	}

	// 判断是否需要生成tls,提前生成secret
	if config.Key != "" && config.Cert != "" {
		if err := c.CreateNetworkSecret(ctx, config); err != nil {
			return err
		}
		isTls = true
	}

	if err := c.CreateNetworkIngress(ctx, config, isTls); err != nil {
		return err
	}

	return nil
}

// DeleteNetwork 清除指定service的网络配置
func (c *Client) DeleteNetwork(ctx *gin.Context, config model.NetworkConfig) error {
	if config.Cert != "" && config.Key != "" {
		if err := c.DeleteNetworkSecret(ctx, config); err != nil {
			return err
		}
	}
	if err := c.DeleteNetworkService(ctx, config); err != nil {
		return err
	}
	return c.DeleteNetworkIngress(ctx, config)
}
