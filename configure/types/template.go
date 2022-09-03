package types

type AllTemplateRequest struct {
	ServiceId int64 `json:"service_id" form:"service_id" binding:"required"`
}

type GetTemplateRequest struct {
	ID    int64 `json:"id" form:"id"`
	SrvId int64 `json:"srv_id" form:"srv_id"`
}

type AddTemplateRequest struct {
	ServiceId   int64  `json:"service_id" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTemplateRequest struct {
	ID int64 `json:"id" form:"id"  binding:"required"`
}
