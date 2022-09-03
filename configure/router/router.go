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
		v1.GET("/environment/filter", handler.AllEnvironmentFilter)
		v1.POST("/environment", handler.AddEnvironment)
		v1.POST("/environment/connect", handler.EnvironmentConnect)
		v1.PUT("/environment", handler.UpdateEnvironment)
		v1.DELETE("/environment", handler.DeleteEnvironment)
		v1.PUT("/environment/service", handler.UpdateEnvService)
		v1.GET("/environment/service", handler.AllEnvService)

		// 服务相关
		v1.GET("/services", handler.AllService)
		v1.GET("/service/envs", handler.AllServiceEnvs)
		v1.POST("/service", handler.AddService)
		v1.PUT("/service", handler.UpdateService)
		v1.DELETE("/service", handler.DeleteService)
		v1.GET("/service/system_field", handler.AllServiceSystemField)
		v1.POST("/service/system_field", handler.AddServiceSystemField)
		v1.GET("/service/fields", handler.AllServiceField)

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

		// 服务字段相关
		v1.GET("/service_field/page", handler.PageServiceField)
		v1.POST("/service_field", handler.AddServiceField)
		v1.PUT("/service_field", handler.UpdateServiceField)
		v1.DELETE("/service_field", handler.DeleteServiceField)
		v1.GET("/service_field_value", handler.AllServiceFieldValue)
		v1.POST("/service_field_value", handler.AddServiceFieldValue)

		// 系统字段相关
		v1.GET("/system_field/page", handler.PageSystemField)
		v1.POST("/system_field", handler.AddSystemField)
		v1.PUT("/system_field", handler.UpdateSystemField)
		v1.DELETE("/system_field", handler.DeleteSystemField)
		v1.GET("/system_field_value", handler.AllSystemFieldValue)
		v1.POST("/system_field_value", handler.AddSystemFieldValue)

		// 操作日志
		v1.GET("operator_log/page", handler.PageOperatorLog)
	}
	return root
}
