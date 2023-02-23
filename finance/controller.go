package finance

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiGroup *gin.RouterGroup) {

	//NetWorth
	apiGroup.POST("/networth", func(ctx *gin.Context) {

		var nw netWorth
		ctx.ShouldBindJSON(&nw)
		result := calculateNetWorth(nw)

		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//CompoundInterest
	apiGroup.POST("/compoundinterest", func(ctx *gin.Context) {

		var ci compoundInterest
		ctx.ShouldBindJSON(&ci)
		result := calculateCompoundInterest(ci)

		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//PriceToEarningsRatio
	apiGroup.POST("/peratio", func(ctx *gin.Context) {

		var pe peRatio
		ctx.ShouldBindJSON(&pe)
		result := calculatePriceToEarningsRatio(pe)

		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//BreakEvenPoint
	apiGroup.POST("/breakevenpoint", func(ctx *gin.Context) {

		var bep breakEvenPoint
		ctx.ShouldBindJSON(&bep)
		result := calculateBreakEvenPoint(bep)

		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//NetIncome
	apiGroup.POST("/netincome", func(ctx *gin.Context) {

		var ni netIncome
		ctx.ShouldBindJSON(&ni)
		result := calculateNetIncome(ni)

		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})

	//CashFlow
	apiGroup.POST("/cashflow", func(ctx *gin.Context) {
		var cf cashFlow
		ctx.ShouldBindJSON(&cf)
		result := calculateCashFlow(cf)

		ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
	})
}
