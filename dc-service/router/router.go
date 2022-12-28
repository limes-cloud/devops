package router

import (
	"dc/handler"
	"github.com/limeschool/gin"
)

func Init() *gin.Engine {
	root := gin.Default()
	root.NoRoute(gin.Resp404())
	root.GET("/healthy", gin.Success())

	v1 := root.Group("/api/v1/").Use(gin.ExtRequestTokenAuth())
	{
		// 网络相关
		v1.GET("/service/pods", handler.GetServicePods)
		v1.GET("/service/release", handler.GetServiceRelease)
		v1.POST("/service", handler.AddService)
		v1.DELETE("/service", handler.DeleteService)
		v1.POST("/network", handler.AddNetwork)
		v1.DELETE("/network", handler.DeleteNetwork)
	}
	return root
}
