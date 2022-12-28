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

		// 服务相关
		v1.GET("/run_types", handler.ServiceRunTypes)
		v1.GET("/variable", handler.AllVariable)
		v1.GET("/service/filter", handler.PageServiceFilter)
		v1.GET("/service/page", handler.PageService)
		v1.GET("/service/envs", handler.AllServiceEnvs)
		v1.POST("/service", handler.AddService)
		v1.PUT("/service", handler.UpdateService)
		v1.DELETE("/service", handler.DeleteService)

		// 代码仓库相关
		v1.GET("/code_registries", handler.AllCodeRegistries)
		v1.GET("/code_registry/filter", handler.AllCodeRegistryFilter)
		v1.GET("/code_registry/types", handler.AllCodeRegistryTypes)
		v1.GET("/code_registry/clone_types", handler.AllCodeRegistryCloneTypes)
		v1.POST("/code_registry", handler.AddCodeRegistry)
		v1.POST("/code_registry/connect", handler.ConnectCodeRegistry)
		v1.PUT("/code_registry", handler.UpdateCodeRegistry)
		v1.DELETE("/code_registry", handler.DeleteCodeRegistry)
		v1.GET("/code_registry/project", handler.GetCodeRegistryProject)
		v1.GET("/code_registry/branches", handler.AllCodeRegistryBranches)
		v1.GET("/code_registry/tags", handler.AllCodeRegistryTags)

		// 镜像仓库相关
		v1.GET("/image_registries", handler.AllImageRegistries)
		v1.GET("/image_registry/filter", handler.AllImageRegistryFilter)
		v1.POST("/image_registry", handler.AddImageRegistry)
		v1.POST("/image_registry/connect", handler.ConnectImageRegistry)
		v1.PUT("/image_registry", handler.UpdateImageRegistry)
		v1.DELETE("/image_registry", handler.DeleteImageRegistry)

		// 打包模板相关
		v1.GET("/dockerfile/page", handler.PageDockerfile)
		v1.GET("/dockerfile/filter", handler.AllDockerfileFilter)
		v1.POST("/dockerfile", handler.AddDockerfile)
		v1.PUT("/dockerfile", handler.UpdateDockerfile)
		v1.DELETE("/dockerfile", handler.DeleteDockerfile)

		// 发布模板相关
		v1.GET("/release/page", handler.PageRelease)
		v1.GET("/release/types", handler.AllReleaseTypes)
		v1.GET("/release/status", handler.AllReleaseStatus)
		v1.POST("/release", handler.AddRelease)
		v1.PUT("/release", handler.UpdateRelease)
		v1.DELETE("/release", handler.DeleteRelease)
		v1.GET("/release/images", handler.AllReleaseImages)

		// 打包相关
		v1.GET("/pack_log/page", handler.PagePackLog)
		v1.POST("/pack", handler.AddPack)

		// 发布相关
		v1.GET("/release_log/page", handler.PageReleaseLog)
		v1.POST("/release_log", handler.AddReleaseLog)

		// 网络相关
		v1.GET("/network/page", handler.PageNetwork)
		v1.POST("/network", handler.AddNetwork)
		v1.PUT("/network", handler.UpdateNetwork)
		v1.DELETE("/network", handler.DeleteNetwork)
	}
	return root
}
