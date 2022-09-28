package types

type PageUserRequest struct {
	Page   int    `json:"page" form:"page" binding:"required" sql:"-"`
	Count  int    `json:"count" form:"count"  binding:"required,max=50"  sql:"-"`
	TeamID int64  `json:"team_id" form:"team_id"`
	RoleID int64  `json:"role_id" form:"role_id"`
	Name   string `json:"name" form:"name" sql:"like '%?%'"`
	Sex    *bool  `json:"sex" form:"sex"`
	Phone  string `json:"phone" form:"phone"`
	Email  string `json:"email" form:"email"`
	Status *bool  `json:"status" form:"status"`
}

type GetUserRequest struct {
	ID int64 `json:"id" form:"id"  binding:"required"`
}

type AddUserRequest struct {
	TeamID    int64  `json:"team_id" binding:"required"`
	RoleID    int64  `json:"role_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Sex       *bool  `json:"sex" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar" binding:"required"`
	Email     string `json:"email"  binding:"required"`
	Status    *bool  `json:"status"  binding:"required"`
	LastLogin int64  `json:"last_login"`
}

type UpdateUserRequest struct {
	ID        int64  `json:"id"  binding:"required"`
	TeamID    int64  `json:"team_id"`
	RoleID    int64  `json:"role_id"`
	Name      string `json:"name"`
	Sex       *bool  `json:"sex"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Status    *bool  `json:"status"`
	Password  string `json:"password"`
	LastLogin int64  `json:"last_login"`
}

type DeleteUserRequest struct {
	ID int64 `json:"id"`
}

type UserLoginRequest struct {
	Phone    string `json:"phone"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Duration     int    `json:"duration"`
}
