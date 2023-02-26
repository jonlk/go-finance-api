package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

func StartService() {
	router := gin.Default()

	registerMiddleware(router)

	apiGroup := router.Group("/api")

	registerHealthCheck(apiGroup)

	finance.RegisterRoutes(apiGroup)

	log.Fatal(router.Run(":80"))
}

func registerHealthCheck(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
}
