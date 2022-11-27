package router

import (
	"configure/handler"
	"github.com/limeschool/gin"
)

func Init() *gin.Engine {
	root := gin.Default()
	root.NoRoute(gin.Resp404())
	root.GET("/healthy", gin.Success())

	v1 := root.Group("/configure")
	{
		// 配置环境相关api
		v1.GET("/environments", handler.AllEnvironment)
		v1.POST("/environment", handler.AddEnvironment)
		v1.POST("/environment/connect", handler.EnvironmentConnect)
		v1.PUT("/environment", handler.UpdateEnvironment)
		v1.DELETE("/environment", handler.DeleteEnvironment)

		// 系统资源
		v1.GET("/resource/page", handler.PageResource)
		v1.POST("/resource", handler.AddResource)
		v1.PUT("/resource", handler.UpdateResource)
		v1.DELETE("/resource", handler.DeleteResource)
		v1.GET("/resource/value", handler.AllResourceValue)
		v1.POST("/resource/value", handler.AddResourceValue)
		v1.GET("/resource/service", handler.AllResourceService)
		v1.POST("/resource/service", handler.AddResourceService)

		// 业务字段相关
		v1.GET("/field/page", handler.PageField)
		v1.POST("/field", handler.AddField)
		v1.PUT("/field", handler.UpdateField)
		v1.DELETE("/field", handler.DeleteField)
		v1.GET("/field_value", handler.AllFieldValue)
		v1.POST("/field_value", handler.AddFieldValue)

		// 获取服务的全部字段以及资源
		v1.GET("/service/fields", handler.AllServiceFieldAndResource)

		// 配置模板相关
		v1.GET("/templates", handler.AllTemplate)
		v1.GET("/template", handler.GetTemplate)
		v1.GET("/template/parse", handler.ParseTemplate)
		v1.POST("/template", handler.AddTemplate)
		v1.PUT("/template", handler.UpdateTemplate)

		// 配置相关
		v1.POST("/config/compare", handler.CompareConfig)
		v1.POST("/config/sync", handler.SyncConfig)
		v1.GET("/config/logs", handler.AllConfigLog)
		v1.GET("/config/log", handler.ConfigLog)
		v1.GET("/config/driver", handler.DriverConfig)
		v1.POST("/config/rollback", handler.RollbackConfig)
		v1.GET("/config", handler.Config)
	}
	return root
}
