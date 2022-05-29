package routes

import (
	"github.com/TCP404/OneTiny-base/core/controller"
	"github.com/TCP404/OneTiny-base/core/middleware"

	"github.com/gin-gonic/gin"
)

func loadCoreRoute(r *gin.Engine) *gin.RouterGroup {
	fileG := r.Group("/file", middleware.CheckLevel)
	{
		fileG.GET("/*filename", controller.Downloader)
		fileG.POST("/upload", controller.Uploader)
	}
	return fileG
}
