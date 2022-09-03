package types

type AddServiceFieldRequest struct {
	ServiceId   int64   `json:"service_id" binding:"required"`
	Field       string  `json:"field" binding:"required"`
	Description *string `json:"description"`
}

type PageServiceFieldRequest struct {
	ServiceId int64  `json:"service_id"  form:"field"`
	Field     string `json:"field" form:"field" sql:"like '%?%'"`
	Page      int    `json:"page" form:"page" sql:"-"`
	Count     int    `json:"count" form:"count" sql:"-"`
}

type UpdateServiceFieldRequest struct {
	ID          int64   `json:"id" binding:"required"`
	Field       string  `json:"field"`
	Value       string  `json:"value"`
	Description *string `json:"description"`
}

type DeleteServiceFieldRequest struct {
	ID int64 `json:"id" binding:"required"`
}
