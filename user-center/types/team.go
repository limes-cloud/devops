package types

type AddTeamRequest struct {
	Name        string `json:"name"  binding:"required"`
	Avatar      string `json:"avatar"  binding:"required"`
	Description string `json:"description"`
	ParentID    int64  `json:"parent_id" binding:"required"`
}

type UpdateTeamRequest struct {
	ID          int64  `json:"id" form:"id" binding:"required"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	ParentID    int64  `json:"parent_id"`
}

type DeleteTeamRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}
