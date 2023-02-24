package web

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(router *gin.Engine) {
	router.Use(gzip.Gzip(gzip.BestSpeed))
}
