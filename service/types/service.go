package types

type AllServiceEnvRequest struct {
	SrvId int64 `json:"srv_id" form:"srv_id"  binding:"required"`
}

type AllServiceFieldRequest struct {
	SrvId int64 `json:"srv_id" form:"srv_id"  binding:"required"`
}

type AddServiceRequest struct {
	Keyword         string  `json:"keyword" binding:"required"`
	Name            string  `json:"name"  binding:"required"`
	IsPrivate       *bool   `json:"is_private"  binding:"required"`
	TeamID          *int64  `json:"team_id"`
	Description     *string `json:"description"`
	EnvIds          []int64 `json:"env_ids" binding:"required"`
	ReleaseID       int64   `json:"release_id" binding:"required"`
	DockerfileID    int64   `json:"dockerfile_id"  binding:"required"`
	CodeRegistryID  int64   `json:"code_registry_id"  binding:"required"`
	ImageRegistryID int64   `json:"image_registry_id" binding:"required"`
	RunType         string  `json:"run_type" binding:"required"`
	RunPort         int64   `json:"run_port" binding:"required"`
	ListenPort      int64   `json:"listen_port" binding:"required"`
	Owner           string  `json:"owner" binding:"required"`
	Repo            string  `json:"repo" binding:"required"`
	Replicas        int64   `json:"replicas" binding:"required"`
	ProbeType       string  `json:"probe_type" binding:"required"`
	ProbeValue      string  `json:"probe_value" binding:"required"`
	ProbeInitDelay  int64   `json:"probe_init_delay" binding:"required"`
}

type PageServiceRequest struct {
	Keyword   string `json:"keyword" form:"keyword"`
	ID        int64  `json:"id" form:"id"`
	Name      string `json:"name" form:"name" sql:"like '%?%'"`
	IsPrivate *bool  `json:"is_private" form:"is_private"`
	TeamID    *int64 `json:"team_id"  form:"team_id" sql:"-"`
	Page      int    `json:"page" form:"page" binding:"required" sql:"-"`
	Count     int    `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
}

type UpdateServiceRequest struct {
	ID              int64   `json:"id" binding:"required"`
	Keyword         string  `json:"keyword"`
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	IsPrivate       *bool   `json:"is_private"`
	TeamID          *int64  `json:"team_id"`
	EnvIds          []int64 `json:"env_ids" binding:"required"`
	ReleaseID       *int64  `json:"release_id"`
	DockerfileID    *int64  `json:"dockerfile_id"`
	CodeRegistryID  *int64  `json:"code_registry_id"`
	ImageRegistryID *int64  `json:"image_registry_id"`
	RunType         *string `json:"run_type" binding:"required"`
	RunPort         *int64  `json:"run_port"`
	ListenPort      *int64  `json:"listen_port"`
	Owner           *string `json:"owner"`
	Repo            *string `json:"repo"`
	Replicas        *int64  `json:"replicas" binding:"required"`
	ProbeType       *string `json:"probe_type" binding:"required"`
	ProbeValue      *string `json:"probe_value" binding:"required"`
	ProbeInitDelay  *int64  `json:"probe_init_delay" binding:"required"`
}

type DeleteServiceRequest struct {
	ID int64 `json:"id" binding:"required"`
}
