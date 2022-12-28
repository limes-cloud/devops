package consts

const (
	RunPort        = "RunPort"
	ListenPort     = "ListenPort"
	Replicas       = "Replicas"
	ProbeValue     = "ProbeValue"
	ProbeInitDelay = "ProbeInitDelay"
	ServiceName    = "ServiceName"
	Image          = "Image"
	Namespace      = "Namespace"

	K8s = "k8s"
	Dc  = "docker-compose"

	K8sLabelPrefix          = "app"
	K8sControllerDeployment = "deployment"
	K8sControllerDaemonSet  = "daemonSet"

	Success = "success"
	Fail    = "fail"
	Check   = "check"
)

var Variables = map[string]string{
	RunPort:        "运行端口",
	ListenPort:     "监听端口",
	Replicas:       "副本数量",
	ProbeValue:     "探测值",
	ProbeInitDelay: "探测延迟时间",
	ServiceName:    "服务标志",
	Image:          "发布镜像",
	Namespace:      "所属环境",
}

var ReleaseStatus = map[string]string{
	Success: "发布成功",
	Fail:    "发布失败",
	Check:   "健康检查",
}
