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

	ums := root.Group("ums/")
	{
		// 角色相关
		ums.GET("/role/data_scope", handler.RoleDataScope)
		ums.GET("/role", handler.AllRole)
		ums.POST("/role", handler.AddRole)
		ums.PUT("/role", handler.UpdateRole)
		ums.DELETE("/role", handler.DeleteRole)
		ums.GET("/role/menu", handler.RoleMenu)
		ums.GET("/role/menu_ids", handler.RoleMenuIds)
		ums.POST("/role/menu", handler.AddRoleMenu)

		// 部门相关
		ums.GET("/team", handler.AllTeam)
		ums.POST("/team", handler.AddTeam)
		ums.PUT("/team", handler.UpdateTeam)
		ums.DELETE("/team", handler.DeleteTeam)

		// 菜单相关
		ums.GET("/menu", handler.AllMenu)
		ums.POST("/menu", handler.AddMenu)
		ums.PUT("/menu", handler.UpdateMenu)
		ums.DELETE("/menu", handler.DeleteMenu)

		// 用户管理相关
		ums.GET("/user/page", handler.PageUser)
		ums.GET("/user", handler.CurUser)
		ums.POST("/user", handler.AddUser)
		ums.POST("/user/login", handler.UserLogin)
		ums.POST("/token/refresh", handler.RefreshToken)
		ums.PUT("/user", handler.UpdateUser)
		ums.DELETE("/user", handler.DeleteUser)

		// 登陆日志
		ums.GET("/login/log", handler.LoginLog)

		v1 := ums.Group("/api/v1", gin.ExtRequestTokenAuth()) //提供给内部访问
		{
			v1.GET("/user", handler.GetUser)
		}
	}
	return root
}
