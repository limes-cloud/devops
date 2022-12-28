package model

const (
	Update         = "发布中"
	RollUpdate     = "滚动升级中"
	RollUpdateFail = "滚动升级失败"
	Available      = "运行中"
	UnAvailable    = "暂停服务"
)

type NetworkConfig struct {
	Namespace   string `json:"namespace"`
	ServiceName string `json:"service_name"`
	Host        string `json:"host"`
	Cert        string `json:"cert"`
	Key         string `json:"key"`
	Redirect    bool   `json:"redirect"`
	TargetPort  int64  `json:"target_port"`
	RunPort     int64  `json:"run_port"`
	Replicas    int64  `json:"replicas"`
}

type ServiceConfig struct {
	Type          string `json:"type"`
	Yaml          string `json:"yaml"`
	Namespace     string `json:"namespace"`
	ServiceName   string `json:"service_name"`
	Replicas      int64  `json:"replicas"`
	RunPort       int64  `json:"run_port" `
	ListenPort    int64  `json:"listen_port"`
	ImageRegistry string `json:"image_registry"`
	ImageUser     string `json:"image_user"`
	ImagePass     string `json:"image_pass"`
}

type Pod struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Image       string `json:"image"`
	Err         string `json:"err"`
	CreatedTime string `json:"created_time"`
	Restart     int    `json:"restart"`
	PodIP       string `json:"pod_ip"`
	Node        string `json:"node"`
}
