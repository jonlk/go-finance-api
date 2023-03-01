package web

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

var ginLambda *ginadapter.GinLambda

func StartService(runAsLambda bool, port int) {

	router := gin.Default()
	registerMiddleware(router)
	apiGroup := router.Group("/api")
	registerHealthCheck(apiGroup)
	finance.RegisterRoutes(apiGroup)

	if runAsLambda {
		ginLambda = ginadapter.New(router)
		lambda.Start(Handler)
	} else {
		log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func registerHealthCheck(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
}
