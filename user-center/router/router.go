package router

import (
	"github.com/limeschool/gin"
	"ums/handler"
)

func Init() *gin.Engine {
	root := gin.Default()
	root.NoRoute(gin.Resp404())
	root.GET("/healthy", gin.Success())
	root.GET("/auth", handler.Auth)

	v1 := root.Group("ums/")
	{
		// 角色相关
		v1.GET("/role/data_scope", handler.RoleDataScope)
		v1.GET("/role", handler.AllRole)
		v1.POST("/role", handler.AddRole)
		v1.PUT("/role", handler.UpdateRole)
		v1.DELETE("/role", handler.DeleteRole)

		// 部门相关
		v1.GET("/team", handler.AllTeam)
		v1.POST("/team", handler.AddTeam)
		v1.PUT("/team", handler.UpdateTeam)
		v1.DELETE("/team", handler.DeleteTeam)
		v1.GET("/role/menu", handler.RoleMenu)
		v1.GET("/role/menu_ids", handler.RoleMenuIds)
		v1.POST("/role/menu", handler.AddRoleMenu)

		// 菜单相关
		v1.GET("/menu", handler.AllMenu)
		v1.POST("/menu", handler.AddMenu)
		v1.PUT("/menu", handler.UpdateMenu)
		v1.DELETE("/menu", handler.DeleteMenu)

		// 用户管理相关
		v1.GET("/user/page", handler.PageUser)
		v1.GET("/user", handler.CurUser)
		v1.POST("/user", handler.AddUser)
		v1.POST("/user/login", handler.UserLogin)
		v1.POST("/token/refresh", handler.RefreshToken)
		v1.PUT("/user", handler.UpdateUser)
		v1.DELETE("/user", handler.DeleteUser)

		// 登陆日志
		v1.GET("/login/log", handler.LoginLog)
	}
	return root
}
