package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/web"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	web.StartService()
}
