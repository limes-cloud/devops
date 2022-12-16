package types

type AddCodeRegistryRequest struct {
	Name      string `json:"name" binding:"required"`
	Desc      string `json:"desc"`
	Type      string `json:"type" binding:"required"`
	Host      string `json:"host" binding:"required"`
	Token     string `json:"token" binding:"required"`
	CloneType string `json:"clone_type"  binding:"required"`
}

type UpdateCodeRegistryRequest struct {
	ID        int64  `json:"id"  binding:"required"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Type      string `json:"type"`
	Host      string `json:"host"`
	Token     string `json:"token"`
	CloneType string `json:"clone_type"`
}

type DeleteCodeRegistryRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type ConnectCodeRegistryRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type GetCodeRegistryProjectRequest struct {
	ID    int64  `form:"id"  json:"id"  binding:"required"`
	Owner string `form:"owner" json:"owner"  binding:"required"`
	Repo  string `form:"repo" json:"repo"  binding:"required"`
}

type AllCodeRegistryBranchesRequest struct {
	ServiceKeyword string `json:"service_keyword" form:"service_keyword" binding:"required"`
}

type AllCodeRegistryTagsRequest struct {
	ServiceKeyword string `json:"service_keyword" form:"service_keyword" binding:"required"`
}
