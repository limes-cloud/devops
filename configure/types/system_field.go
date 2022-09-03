package types

type AddSystemFieldRequest struct {
	Field       string  `json:"field" binding:"required"`
	Type        string  `json:"type"  binding:"required"`
	ChildField  string  `json:"child_field" binding:"required"`
	Description *string `json:"description"`
}

type PageSystemFieldRequest struct {
	Field string `json:"field" form:"field" sql:"like '%?%'"`
	Type  string `json:"type" form:"type"`
	Page  int    `json:"page" form:"page" sql:"-"`
	Count int    `json:"count" form:"count" sql:"-"`
}

type UpdateSystemFieldRequest struct {
	ID          int64   `json:"id" binding:"required"`
	Field       string  `json:"field"`
	ChildField  string  `json:"child_field"`
	Type        string  `json:"type"  binding:"required"`
	Description *string `json:"description"`
}

type DeleteSystemFieldRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type AllServiceSystemFieldRequest struct {
	ServiceId int64 `json:"service_id" form:"service_id"  binding:"required"`
}

type AddServiceSystemFieldRequest struct {
	FieldIds  []int64 `json:"field_ids" binding:"required"`
	ServiceId int64   `json:"service_id"  binding:"required"`
}
