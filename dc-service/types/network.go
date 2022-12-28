package types

type AddNetworkRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
	RunPort     int    `json:"run_port" binding:"required"`
	Replicas    int    `json:"replicas" binding:"required"`
	Host        string `json:"host" binding:"required"`
	Cert        string `json:"cert"`
	Key         string `json:"key"`
	Redirect    bool   `json:"redirect"`
}

type DeleteNetworkRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
	Host        string `json:"host" binding:"required"`
}
