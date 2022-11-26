package types

type AddResourceRequest struct {
	Field       string  `json:"field" binding:"required"`
	Type        string  `json:"type"  binding:"required"`
	ChildField  string  `json:"child_field" binding:"required"`
	Description *string `json:"description"`
}

type PageResourceRequest struct {
	Field string `json:"field" form:"field" sql:"like '%?%'"`
	Type  string `json:"type" form:"type"`
	Page  int    `json:"page" form:"page" sql:"-"`
	Count int    `json:"count" form:"count" sql:"-"`
}

type UpdateResourceRequest struct {
	ID          int64   `json:"id" binding:"required"`
	Field       string  `json:"field"`
	ChildField  string  `json:"child_field"`
	Type        string  `json:"type"  binding:"required"`
	Description *string `json:"description"`
}

type DeleteResourceRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type AllServiceResourceRequest struct {
	ServiceId int64 `json:"service_id" form:"service_id"  binding:"required"`
}

type AddServiceResourceRequest struct {
	FieldIds  []int64 `json:"field_ids" binding:"required"`
	ServiceId int64   `json:"service_id"  binding:"required"`
}
