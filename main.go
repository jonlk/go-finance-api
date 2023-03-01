package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/web"
)

const (
	DEVELOPMENT_PORT int16 = 3000
	RELEASE_PORT     int16 = 80
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-r" {
		gin.SetMode(gin.ReleaseMode)
		web.StartService(RELEASE_PORT)
	} else {
		web.StartService(DEVELOPMENT_PORT)
	}
}
