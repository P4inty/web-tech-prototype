package main

import (
	"webtech/prototype/file"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	r.POST("/upload/meta", file.UploadMeta)
	r.POST("/upload/:uri", file.Upload)
	r.GET("/uploads", file.All)
	r.GET("/download/:uri", file.Download)
}
