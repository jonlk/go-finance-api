package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

func StartService() {

	router := gin.Default()

	registerMiddleware(router)

	apiGroup := router.Group("/api")

	finance.RegisterRoutes(apiGroup)

	log.Fatal(router.Run(":3000"))
}
