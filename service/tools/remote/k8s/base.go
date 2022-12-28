package k8s

import (
	"crypto/tls"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net"
	"net/http"
	"service/consts"
	"time"
)

type Client struct {
	client *kubernetes.Clientset
}

func NewK8sClient(host, token string) (*Client, error) {
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

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

// objectMeta 获取meta
func (c *Client) objectMeta(namespace, srv string) meta.ObjectMeta {
	return meta.ObjectMeta{
		Name:      srv,
		Namespace: namespace,
		Labels: map[string]string{
			"from":                "devops",
			consts.K8sLabelPrefix: srv,
		},
	}
}

// selector 获取选择器
func (c *Client) selector(srv string) map[string]string {
	return map[string]string{
		consts.K8sLabelPrefix: srv,
	}
}
