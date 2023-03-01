package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

func StartService(port int16) {
	router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.1.2"})
	registerMiddleware(router)
	apiGroup := router.Group("/api")
	registerHealthCheck(apiGroup)
	finance.RegisterRoutes(apiGroup)
	log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}

func registerHealthCheck(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
}
