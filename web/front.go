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

	registerRoute("/basicliquidityratio", &finance.BasicLiquidityRatio{}, apiGroup)
	registerRoute("/breakevenpoint", &finance.BreakEvenPoint{}, apiGroup)
	registerRoute("/cashflow", &finance.CashFlow{}, apiGroup)
	registerRoute("/compoundinterest", &finance.CompoundInterest{}, apiGroup)
	registerRoute("/netincome", &finance.NetIncome{}, apiGroup)
	registerRoute("/networth", &finance.NetWorth{}, apiGroup)
	registerRoute("/peratio", &finance.PERatio{}, apiGroup)
	registerRoute("/ruleof72", &finance.RuleOf72{}, apiGroup)
	registerRoute("/simpleinterest", &finance.SimpleInterest{}, apiGroup)
	registerRoute("/variationofinvestment", &finance.VariationOfInvestment{}, apiGroup)

	if runAsLambda {
		ginLambda = ginadapter.New(router)
		lambda.Start(handler)
	} else {
		log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
	}
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func registerHealthCheck(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
}

func registerRoute(route string, calculation finance.Calculation, apiGroup *gin.RouterGroup) {
	apiGroup.GET(route, func(ctx *gin.Context) {
		if err := ctx.ShouldBind(calculation); err != nil {
			ctx.Error(err)
		} else {
			result := finance.GetFinanceResult(calculation)
			buildResponseModel(&result, ctx)
		}
	})
}
