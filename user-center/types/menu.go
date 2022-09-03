package types

type AllMenuRequest struct {
	IsFilter *bool `json:"is_filter" form:"is_filter"`
}

type AddMenuRequest struct {
	Title      string `json:"title" binding:"required"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Type       string `json:"type" binding:"required"`
	Permission string `json:"permission"`
	Method     string `json:"method"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	ParentID   int64  `json:"parent_id" binding:"required"`
	Weight     int    `json:"weight"`
	Hidden     bool   `json:"hidden"`
	IsFrame    bool   `json:"is_frame"`
}

type UpdateMenuRequest struct {
	ID         int64  `json:"id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Type       string `json:"type" binding:"required"`
	Permission string `json:"permission"`
	Method     string `json:"method"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	ParentID   int64  `json:"parent_id" binding:"required"`
	Weight     *int   `json:"weight"`
	Hidden     *bool  `json:"hidden"`
	IsFrame    *bool  `json:"is_frame"`
}

type DeleteMenuRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}
