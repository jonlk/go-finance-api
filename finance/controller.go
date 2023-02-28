package finance

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func calculationResponse(c calculation, ctx *gin.Context) {
	result := c.calculate()
	ctx.String(http.StatusOK, fmt.Sprintf("%.2f", result))
}

func RegisterRoutes(apiGroup *gin.RouterGroup) {

	apiGroup.GET("/basicliquidityratio", func(ctx *gin.Context) {
		var blr basicLiquidityRatio
		if err := ctx.Bind(&blr); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(blr, ctx)
		}
	})

	apiGroup.GET("/breakevenpoint", func(ctx *gin.Context) {
		var bep breakEvenPoint
		if err := ctx.Bind(&bep); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(bep, ctx)
		}
	})

	apiGroup.GET("/cashflow", func(ctx *gin.Context) {
		var cf cashFlow
		if err := ctx.Bind(&cf); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(cf, ctx)
		}
	})

	apiGroup.GET("/compoundinterest", func(ctx *gin.Context) {
		var ci compoundInterest
		if err := ctx.Bind(&ci); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(ci, ctx)
		}
	})

	apiGroup.GET("/netincome", func(ctx *gin.Context) {
		var ni netIncome
		if err := ctx.Bind(&ni); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(ni, ctx)
		}
	})

	apiGroup.GET("/networth", func(ctx *gin.Context) {
		var nw netWorth
		if err := ctx.Bind(&nw); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(nw, ctx)
		}
	})

	apiGroup.GET("/peratio", func(ctx *gin.Context) {
		var pe peRatio
		if err := ctx.Bind(&pe); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(pe, ctx)
		}
	})

	apiGroup.GET("/ruleof72", func(ctx *gin.Context) {
		var cir compoundInterestRate
		if err := ctx.Bind(&cir); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(cir, ctx)
		}
	})

	apiGroup.GET("/simpleinterest", func(ctx *gin.Context) {
		var si simpleInterest
		if err := ctx.Bind(&si); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(si, ctx)
		}
	})

	apiGroup.GET("/variationofinvestment", func(ctx *gin.Context) {
		var vi variationOfInvestment
		if err := ctx.Bind(&vi); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			calculationResponse(vi, ctx)
		}
	})
}
