package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/web"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "-r" {
		gin.SetMode(gin.ReleaseMode)
	}

	web.StartService()
}
