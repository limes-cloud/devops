package types

type GetNoticeRequest struct {
	Page   int    `json:"page" form:"page" sql:"-"`
	Count  int    `json:"count" form:"count" sql:"-"`
	Title  string `json:"title" form:"title" sql:"like '%?%'"`
	Status *bool  `json:"status" form:"status"`
}

type AddNoticeRequest struct {
	Cid        string  `json:"cid" binding:"required"`
	Title      string  `json:"title" binding:"required"`
	Rule       string  `json:"rule" binding:"required"`
	UserIds    string  `json:"user_ids" binding:"required"`
	ChannelIds []int64 `json:"channel_ids"  binding:"required"`
	Value      int64   `json:"value"`
}

type UpdateNoticeRequest struct {
	ID         int64   `json:"id" binding:"required"`
	Cid        string  `json:"cid"`
	Title      string  `json:"title"`
	Rule       string  `json:"rule"`
	UserIds    string  `json:"user_ids"`
	Status     *bool   `json:"status"`
	ChannelIds []int64 `json:"channel_ids"`
}

type DeleteNoticeRequest struct {
	ID int64 `json:"id" binding:"required"`
}
