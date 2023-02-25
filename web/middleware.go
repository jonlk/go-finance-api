package web

import (
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func registerMiddleware(router *gin.Engine) {

	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(customMiddleware)
	router.Use(errorHandler)

}

var customMiddleware = func(ctx *gin.Context) {
	//EXAMPLE: add a custom return header
	ctx.Header("Custom-Header", "Returned to client")
	ctx.Next()
}

var errorHandler = func(ctx *gin.Context) {
	for _, err := range ctx.Errors {
		log.Print(map[string]any{
			"Err":  err.Error(),
			"Type": err.Type,
			"Meta": err.Meta,
		})
	}
	ctx.Next()
}
