package types

type AllEnvironmentRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
	Name    string `json:"name" form:"name" sql:"like '%?%'"`
	Status  *bool  `json:"status" form:"status"`
}

type EnvironmentConnectRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type AddEnvironmentRequest struct {
	Keyword      string `json:"keyword" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Type         string `json:"type" binding:"required"`
	DcHost       string `json:"dc_host"`
	DcToken      string `json:"dc_token"`
	K8sHost      string `json:"k8s_host"`
	K8sToken     string `json:"k8s_token"`
	K8sNamespace string `json:"k8s_namespace"`
	Status       bool   `json:"status"`
}

type UpdateEnvironmentRequest struct {
	ID           int64   `json:"id"  binding:"required"`
	Name         *string `json:"name"`
	Keyword      *string `json:"keyword"`
	Type         string  `json:"type" binding:"required"`
	Description  *string `json:"description"`
	Status       *bool   `json:"status"`
	DcHost       *string `json:"dc_host"`
	DcToken      *string `json:"dc_token"`
	K8sHost      *string `json:"k8s_host"`
	K8sToken     *string `json:"k8s_token"`
	K8sNamespace *string `json:"k8s_namespace"`
}

type DeleteEnvironmentRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type UpdateServiceEnvRequest struct {
	ID     int64   `json:"id"  binding:"required"`
	SrvIds []int64 `json:"srv_ids" binding:"required"`
}
