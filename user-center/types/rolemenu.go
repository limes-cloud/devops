package types

type AddRoleMenuRequest struct {
	RoleID  int64   `json:"role_id"`
	MenuIds []int64 `json:"menu_ids"`
}

type RoleMenuIdsRequest struct {
	RoleID int64 `json:"role_id" form:"role_id"`
}
