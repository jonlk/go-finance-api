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
		panic("could not determine operating mode")
	}

	web.StartService(runAsLambda, PORT)
}
