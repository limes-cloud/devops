package image_release

import (
	"github.com/limeschool/gin"
)

type ImageRelease interface {
	DeleteFormYaml(ctx *gin.Context, applyYaml string) (err error)
	UpdateFromYaml(ctx *gin.Context, applyYaml string) (err error)
	GetStartStatus(ctx *gin.Context, srv string) error
	AddNetwork(ctx *gin.Context, config NetworkConfig) (err error)
	DeleteNetwork(ctx *gin.Context, config NetworkConfig) (err error)
}

type NetworkConfig struct {
	Namespace  string `json:"namespace"`
	Service    string `json:"service"`
	Host       string `json:"host"`
	Cert       string `json:"cert"`
	Key        string `json:"key"`
	Redirect   bool   `json:"redirect"`
	TargetPort int64  `json:"target_port"`
	RunPort    int64  `json:"run_port"`
}
