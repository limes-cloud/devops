package router

import (
	"github.com/limeschool/gin"
	"service/handler"
)

func Init() *gin.Engine {
	root := gin.Default()
	root.NoRoute(gin.Resp404())
	root.GET("/healthy", gin.Success())

	v1 := root.Group("/service")
	{
		// 配置环境相关api
		v1.GET("/environments", handler.AllEnvironment)
		v1.GET("/environment/filter", handler.AllEnvironmentFilter)
		v1.POST("/environment", handler.AddEnvironment)
		v1.PUT("/environment", handler.UpdateEnvironment)
		v1.DELETE("/environment", handler.DeleteEnvironment)

		// 服务环境所属相关
		//v1.PUT("/environment/service", handler.UpdateEnvService)
		//v1.GET("/environment/service", handler.AllEnvService)

		// 服务相关
		v1.GET("/service/page", handler.PageService)
		v1.GET("/service/envs", handler.AllServiceEnvs)
		v1.POST("/service", handler.AddService)
		v1.PUT("/service", handler.UpdateService)
		v1.DELETE("/service", handler.DeleteService)

	}
	return root
}
