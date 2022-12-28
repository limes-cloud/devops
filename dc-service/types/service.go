package types

type AddServiceRequest struct {
	Yaml          string `json:"yaml" binding:"required"`
	ServiceName   string `json:"service_name" binding:"required"`
	RunPort       int    `json:"run_port" binding:"required"`
	ListenPort    int    `json:"listen_port" binding:"required"`
	Replicas      int    `json:"replicas" binding:"required"`
	ImageRegistry string `json:"image_registry"`
	ImageUser     string `json:"image_user"`
	ImagePass     string `json:"image_pass"`
}

type DeleteServiceRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
	Replicas    int    `json:"replicas" binding:"required"`
}

type GetServiceReleaseRequest struct {
	ServiceName string `form:"service_name" binding:"required"`
	Replicas    int    `form:"replicas" binding:"required"`
}

type GetServicePodsRequest struct {
	ServiceName string `form:"service_name" binding:"required"`
	Replicas    int    `form:"replicas" binding:"required"`
}
