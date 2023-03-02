package main

import (
	"os"
	"strconv"

	"github.com/jonlk/go-finance-api/web"
)

var (
	PORT int = 3000
)

func main() {
	runAsLambda, err := strconv.ParseBool(os.Getenv("RUN_AS_LAMBDA"))
	if err != nil {
		panic("RUN_AS_LAMBDA environment variable not set (true or false)")
	}

	web.StartService(runAsLambda, PORT)
}
