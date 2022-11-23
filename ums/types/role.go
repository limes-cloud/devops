package types

type AllRoleRequest struct {
	Name    string `json:"name" form:"name" sql:"like '%?%'"`
	Keyword string `json:"keyword" form:"keyword"`
	Status  *bool  `json:"status" form:"status"`
}

type AddRoleRequest struct {
	Name        string `json:"name"  binding:"required"`
	Keyword     string `json:"keyword" binding:"required"`
	Status      *bool  `json:"status"  binding:"required"`
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	TeamIds     string `json:"team_ids"`
	DataScope   string `json:"data_scope"  binding:"required"`
}

type UpdateRoleRequest struct {
	ID          int64  `json:"id" binding:"required"`
	Name        string `json:"name"`
	Status      *bool  `json:"status"`
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	DataScope   string `json:"data_scope"`
	TeamIds     string `json:"team_ids"`
}

type DeleteRoleRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}
