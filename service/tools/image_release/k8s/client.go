package k8s

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/limeschool/gin"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	syaml "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"net"
	"net/http"
	sigyaml "sigs.k8s.io/yaml"
	"time"
)

type Client struct {
	KubernetesClient *kubernetes.Clientset
	DynamicClient    dynamic.Interface
	Namespace        string
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
	namespace := c.Namespace
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
		gvr, err := c.GtGVR(object.GroupVersionKind())
		if err != nil {
			return err
		}

		if err != nil {
			return fmt.Errorf("unable to marshal resource as yaml: %w", err)
		}

		// 当为设置namespace 以yaml中的为准
		if namespace == "" {
			namespace = object.GetNamespace()
		}

		namespaceRes := c.DynamicClient.Resource(gvr).Namespace(namespace)
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
	namespace := c.Namespace

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
		gvr, err := c.GtGVR(object.GroupVersionKind())
		if err != nil {
			return err
		}
		unstructuredYaml, err := sigyaml.Marshal(object)
		if err != nil {
			return fmt.Errorf("unable to marshal resource as yaml: %w", err)
		}

		// 当为设置namespace 以yaml中的为准
		if namespace == "" {
			namespace = object.GetNamespace()
		}

		namespaceRes := c.DynamicClient.Resource(gvr).Namespace(namespace)
		// 获取不到资源则进行创建资源
		if _, getErr := namespaceRes.Get(ctx, object.GetName(), v1.GetOptions{}); getErr != nil {
			_, err = namespaceRes.Create(ctx, object, v1.CreateOptions{})
			return err
		}

		// 获取到资源则进行修改
		force := true
		if _, err = namespaceRes.Patch(ctx, object.GetName(), types.ApplyPatchType, unstructuredYaml, v1.PatchOptions{
			FieldManager: object.GetName(),
			Force:        &force,
		}); err != nil {
			return fmt.Errorf("unable to patch resource: %w", err)
		}
	}
	return nil
}
