package types

type AddFieldRequest struct {
	ServiceKeyword string  `json:"service_keyword" binding:"required"`
	Field          string  `json:"field" binding:"required"`
	Description    *string `json:"description"`
}

type PageFieldRequest struct {
	ServiceKeyword string `json:"service_keyword"  form:"field"`
	Field          string `json:"field" form:"field" sql:"like '%?%'"`
	Page           int    `json:"page" form:"page" sql:"-"`
	Count          int    `json:"count" form:"count" sql:"-"`
}

type UpdateFieldRequest struct {
	ID          int64   `json:"id" binding:"required"`
	Field       string  `json:"field"`
	Value       string  `json:"value"`
	Description *string `json:"description"`
}

type DeleteFieldRequest struct {
	ID int64 `json:"id" binding:"required"`
}
