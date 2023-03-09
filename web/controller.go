package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jonlk/go-finance-api/finance"
)

func registerRoutes(apiGroup *gin.RouterGroup) {

	apiGroup.GET("/basicliquidityratio", func(ctx *gin.Context) {
		var blr finance.BasicLiquidityRatio
		if err := ctx.ShouldBind(&blr); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&blr, ctx)
		}
	})

	apiGroup.GET("/breakevenpoint", func(ctx *gin.Context) {
		var bep finance.BreakEvenPoint
		if err := ctx.ShouldBind(&bep); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&bep, ctx)
		}
	})

	apiGroup.GET("/cashflow", func(ctx *gin.Context) {
		var cf finance.CashFlow
		if err := ctx.ShouldBind(&cf); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&cf, ctx)
		}
	})

	apiGroup.GET("/compoundinterest", func(ctx *gin.Context) {
		var ci finance.CompoundInterest
		if err := ctx.ShouldBind(&ci); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&ci, ctx)
		}
	})

	apiGroup.GET("/netincome", func(ctx *gin.Context) {
		var ni finance.NetIncome
		if err := ctx.ShouldBind(&ni); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&ni, ctx)
		}
	})

	apiGroup.GET("/networth", func(ctx *gin.Context) {
		var nw finance.NetWorth
		if err := ctx.ShouldBind(&nw); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&nw, ctx)
		}
	})

	apiGroup.GET("/peratio", func(ctx *gin.Context) {
		var pe finance.PERatio
		if err := ctx.ShouldBind(&pe); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&pe, ctx)
		}
	})

	apiGroup.GET("/ruleof72", func(ctx *gin.Context) {
		var cir finance.RuleOf72
		if err := ctx.ShouldBind(&cir); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&cir, ctx)
		}
	})

	apiGroup.GET("/simpleinterest", func(ctx *gin.Context) {
		var si finance.SimpleInterest
		if err := ctx.ShouldBind(&si); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&si, ctx)
		}
	})

	apiGroup.GET("/variationofinvestment", func(ctx *gin.Context) {
		var vi finance.VariationOfInvestment
		if err := ctx.ShouldBind(&vi); err != nil {
			ctx.Error(err)
		} else {
			finance.BuildCalculationResponse(&vi, ctx)
		}
	})
}
