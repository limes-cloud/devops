package remote

import (
	"github.com/limeschool/gin"
	"service/consts"
	"service/errors"
	"service/tools/remote/dc"
	"service/tools/remote/k8s"
	"service/tools/remote/model"
)

type Remote interface {
	DeleteService(ctx *gin.Context, srv model.ServiceConfig) error
	CreateService(ctx *gin.Context, srv model.ServiceConfig) error
	GetServiceRelease(ctx *gin.Context, srv model.ServiceConfig) error
	GetServicePods(ctx *gin.Context, srv model.ServiceConfig) ([]model.Pod, error)
	CreateNetwork(ctx *gin.Context, config model.NetworkConfig) error
	DeleteNetwork(ctx *gin.Context, config model.NetworkConfig) error
}

func NewClient(tp, host, token string) (Remote, error) {
	if tp == consts.K8s {
		return k8s.NewK8sClient(host, token)
	}
	if tp == consts.Dc {
		return dc.NewClient(host, token)
	}
	return nil, errors.New("无法识别的类型")
}
