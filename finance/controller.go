package finance

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func calculationResponse(c calculation, ctx *gin.Context) {
	result := c.calculate()
	ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
}

func RegisterRoutes(apiGroup *gin.RouterGroup) {

	apiGroup.POST("/basicliquidityratio", func(ctx *gin.Context) {
		var blr basicLiquidityRatio
		ctx.ShouldBindJSON(&blr)
		calculationResponse(blr, ctx)
	})

	apiGroup.POST("/breakevenpoint", func(ctx *gin.Context) {
		var bep breakEvenPoint
		ctx.ShouldBindJSON(&bep)
		calculationResponse(bep, ctx)
	})

	apiGroup.POST("/cashflow", func(ctx *gin.Context) {
		var cf cashFlow
		ctx.ShouldBindJSON(&cf)
		calculationResponse(cf, ctx)
	})

	apiGroup.POST("/compoundinterest", func(ctx *gin.Context) {
		var ci compoundInterest
		ctx.ShouldBindJSON(&ci)
		calculationResponse(ci, ctx)
	})

	apiGroup.POST("/netincome", func(ctx *gin.Context) {
		var ni netIncome
		ctx.ShouldBindJSON(&ni)
		calculationResponse(ni, ctx)
	})

	apiGroup.POST("/networth", func(ctx *gin.Context) {
		var nw netWorth
		ctx.ShouldBindJSON(&nw)
		calculationResponse(nw, ctx)
	})

	apiGroup.POST("/peratio", func(ctx *gin.Context) {
		var pe peRatio
		ctx.ShouldBindJSON(&pe)
		calculationResponse(pe, ctx)
	})

	apiGroup.GET("/ruleof72/:cir", func(ctx *gin.Context) {
		input := ctx.Param("cir")
		if value, err := strconv.ParseFloat(input, 64); err == nil {
			calculationResponse(compoundInterestRate(value), ctx)
		} else {
			ctx.String(http.StatusBadRequest, "Could not parse float")
		}
	})

	apiGroup.POST("/simpleinterest", func(ctx *gin.Context) {
		var si simpleInterest
		ctx.ShouldBindJSON(&si)
		calculationResponse(si, ctx)
	})

	apiGroup.POST("/variationofinvestment", func(ctx *gin.Context) {
		var vi variationOfInvestment
		ctx.ShouldBindJSON(&vi)
		calculationResponse(vi, ctx)
	})
}
