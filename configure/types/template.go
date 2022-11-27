package types

type AllTemplateRequest struct {
	ServiceKeyword string `json:"service_keyword" form:"service_keyword" binding:"required"`
}

type GetTemplateRequest struct {
	ID      int64  `json:"id" form:"id"`
	Keyword string `json:"keyword" form:"keyword"`
}

type AddTemplateRequest struct {
	ServiceKeyword string `json:"service_keyword" binding:"required"`
	Content        string `json:"content" binding:"required"`
	Description    string `json:"description" binding:"required"`
}

type UpdateTemplateRequest struct {
	ID int64 `json:"id" form:"id"  binding:"required"`
}

type ParseTemplateRequest struct {
	SrvKeyword string `json:"srv_keyword" form:"srv_keyword" binding:"required"`
	EnvKeyword string `json:"env_keyword" form:"env_keyword" binding:"required"`
}
