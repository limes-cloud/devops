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

type AllResourceServiceRequest struct {
	ResourceID int64 `json:"resource_id" form:"resource_id"  binding:"required"`
}

type AddResourceServiceRequest struct {
	ServiceKeywords []string `json:"service_keywords" binding:"required"`
	ResourceID      int64    `json:"resource_id"  binding:"required"`
}
