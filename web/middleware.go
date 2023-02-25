package web

import (
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func registerMiddleware(router *gin.Engine) {
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(exampleCustomMiddleware)
	router.Use(errorHandler)
}

var exampleCustomMiddleware = func(ctx *gin.Context) {

	//Retrieve a header on request
	interceptedHeader := ctx.GetHeader("Content-Type")
	log.Printf("The Content-Type header is: %v\n", interceptedHeader)

	//Or add a custom return header in response
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
