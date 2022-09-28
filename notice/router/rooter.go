package router

import (
	"github.com/limeschool/gin"
	"notice/handler"
)

func Init() *gin.Engine {
	root := gin.Default()
	root.NoRoute(gin.Resp404())
	root.GET("/healthy", gin.Success())
	notice := root.Group("/notice")
	{
		notice.GET("/channels", handler.AllChannel)
		notice.GET("/channel/filter", handler.AllChannelFilter)
		notice.POST("/channel", handler.AddChannel)
		notice.PUT("/channel", handler.UpdateChannel)
		notice.DELETE("/channel", handler.DeleteChannel)

		notice.GET("/notice/page", handler.PageNotice)
		notice.POST("/notice", handler.AddNotice)
		notice.PUT("/notice", handler.UpdateNotice)
		notice.DELETE("/notice", handler.DeleteNotice)

		notice.GET("/log/page", handler.PageLog)
	}
	return root
}
