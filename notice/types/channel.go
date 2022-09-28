package types

type GetChannelRequest struct {
	Name   string `json:"name" form:"name" sql:"like '%?%'"`
	Status *bool  `json:"status" form:"status"`
}

type AddChannelRequest struct {
	Name   string `json:"name" binding:"required"`
	Config string `json:"config" binding:"required"`
	Status *bool  `json:"status" binding:"required"`
}

type UpdateChannelRequest struct {
	ID     int64  `json:"id"  binding:"required"`
	Name   string `json:"name"`
	Config string `json:"config"`
	Status *bool  `json:"status"`
}

type DeleteChannelRequest struct {
	ID int64 `json:"id"  binding:"required"`
}
