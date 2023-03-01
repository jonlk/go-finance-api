package web

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

var ginLambda *ginadapter.GinLambda

func StartService() {
	router := gin.Default()
	registerMiddleware(router)
	apiGroup := router.Group("/api")
	registerHealthCheck(apiGroup)
	finance.RegisterRoutes(apiGroup)

	ginLambda = ginadapter.New(router)

	lambda.Start(Handler)

	//log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func registerHealthCheck(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
}
