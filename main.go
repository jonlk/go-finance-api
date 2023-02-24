package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "-r" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	apiGroup := router.Group("/api")
	finance.RegisterRoutes(apiGroup)
	log.Fatal(router.Run(":3000"))
}
