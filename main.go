package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	apiGroup := router.Group("/api")
	finance.RegisterRoutes(apiGroup)
	log.Fatal(router.Run(":3000"))
}
