package web

import (
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(router *gin.Engine) {
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(errorHandler)
}

var errorHandler = func(ctx *gin.Context) {

	ctx.Next()

	for _, err := range ctx.Errors {
		log.Print(map[string]any{
			"Err":  err.Error(),
			"Type": err.Type,
			"Meta": err.Meta,
		})
	}
}
