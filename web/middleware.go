package web

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func registerMiddleware(router *gin.Engine) {
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(errorHandler)
}

var errorHandler = func(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		for _, err := range ctx.Errors {
			log.Print(map[string]any{
				"Err":  err.Error(),
				"Type": err.Type,
				"Meta": err.Meta,
			})
		}
		ctx.String(http.StatusBadRequest, strings.Join(ctx.Errors.Errors(), ","))
	}
}
