package finance

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiGroup *gin.RouterGroup) {

	apiGroup.GET("/basicliquidityratio", func(ctx *gin.Context) {
		var blr basicLiquidityRatio
		if err := ctx.ShouldBind(&blr); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&blr, ctx)
		}
	})

	apiGroup.GET("/breakevenpoint", func(ctx *gin.Context) {
		var bep breakEvenPoint
		if err := ctx.ShouldBind(&bep); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&bep, ctx)
		}
	})

	apiGroup.GET("/cashflow", func(ctx *gin.Context) {
		var cf cashFlow
		if err := ctx.ShouldBind(&cf); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&cf, ctx)
		}
	})

	apiGroup.GET("/compoundinterest", func(ctx *gin.Context) {
		var ci compoundInterest
		if err := ctx.ShouldBind(&ci); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&ci, ctx)
		}
	})

	apiGroup.GET("/netincome", func(ctx *gin.Context) {
		var ni netIncome
		if err := ctx.ShouldBind(&ni); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&ni, ctx)
		}
	})

	apiGroup.GET("/networth", func(ctx *gin.Context) {
		var nw netWorth
		if err := ctx.ShouldBind(&nw); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&nw, ctx)
		}
	})

	apiGroup.GET("/peratio", func(ctx *gin.Context) {
		var pe peRatio
		if err := ctx.ShouldBind(&pe); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&pe, ctx)
		}
	})

	apiGroup.GET("/ruleof72", func(ctx *gin.Context) {
		var cir ruleOf72
		if err := ctx.ShouldBind(&cir); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&cir, ctx)
		}
	})

	apiGroup.GET("/simpleinterest", func(ctx *gin.Context) {
		var si simpleInterest
		if err := ctx.ShouldBind(&si); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&si, ctx)
		}
	})

	apiGroup.GET("/variationofinvestment", func(ctx *gin.Context) {
		var vi variationOfInvestment
		if err := ctx.ShouldBind(&vi); err != nil {
			ctx.Error(err)
		} else {
			buildCalculationResponse(&vi, ctx)
		}
	})
}
