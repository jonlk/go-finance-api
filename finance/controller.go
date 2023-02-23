package finance

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiGroup *gin.RouterGroup) {

	//calculateNetWorth
	apiGroup.POST("/networth", func(ctx *gin.Context) {
		var netWorth netWorth
		ctx.ShouldBindJSON(&netWorth)
		result := calculateNetWorth(netWorth)
		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//calculateCompoundInterest
	apiGroup.POST("/compoundinterest", func(ctx *gin.Context) {
		var compoundInterest compoundInterest
		ctx.ShouldBindJSON(&compoundInterest)
		result := calculateCompoundInterest(compoundInterest)
		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//calculatePriceToEarningsRatio
	apiGroup.POST("/peratio", func(ctx *gin.Context) {
		var peRatio peRatio
		ctx.ShouldBindJSON(&peRatio)
		result := calculatePriceToEarningsRatio(peRatio)
		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//calculateBreakEvenPoint
	apiGroup.POST("/breakevenpoint", func(ctx *gin.Context) {
		var breakEvenPoint breakEvenPoint
		ctx.ShouldBindJSON(&breakEvenPoint)
		result := calculateBreakEvenPoint(breakEvenPoint)
		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//calculateNetIncome
	apiGroup.POST("/netincome", func(ctx *gin.Context) {
		var netIncome netIncome
		ctx.ShouldBindJSON(&netIncome)
		result := calculateNetIncome(netIncome)
		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})
}
