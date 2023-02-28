package finance

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type calculationResponse struct {
	ResponseId      uuid.UUID `json:"responseId"`
	Timestamp       int64     `json:"timestamp"`
	CalculationType string    `json:"calculationType"`
	Value           float64   `json:"value"`
}

func buildCalculationResponse(c calculation, ctx *gin.Context) {
	value := c.calculate()
	response := calculationResponse{
		ResponseId:      uuid.New(),
		Timestamp:       time.Now().Unix(),
		CalculationType: reflect.TypeOf(c).Name(),
		Value:           value,
	}
	ctx.JSON(http.StatusOK, &response)
}

func RegisterRoutes(apiGroup *gin.RouterGroup) {

	apiGroup.GET("/basicliquidityratio", func(ctx *gin.Context) {
		var blr basicLiquidityRatio
		if err := ctx.Bind(&blr); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(blr, ctx)
		}
	})

	apiGroup.GET("/breakevenpoint", func(ctx *gin.Context) {
		var bep breakEvenPoint
		if err := ctx.Bind(&bep); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(bep, ctx)
		}
	})

	apiGroup.GET("/cashflow", func(ctx *gin.Context) {
		var cf cashFlow
		if err := ctx.Bind(&cf); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(cf, ctx)
		}
	})

	apiGroup.GET("/compoundinterest", func(ctx *gin.Context) {
		var ci compoundInterest
		if err := ctx.Bind(&ci); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(ci, ctx)
		}
	})

	apiGroup.GET("/netincome", func(ctx *gin.Context) {
		var ni netIncome
		if err := ctx.Bind(&ni); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(ni, ctx)
		}
	})

	apiGroup.GET("/networth", func(ctx *gin.Context) {
		var nw netWorth
		if err := ctx.Bind(&nw); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(nw, ctx)
		}
	})

	apiGroup.GET("/peratio", func(ctx *gin.Context) {
		var pe peRatio
		if err := ctx.Bind(&pe); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(pe, ctx)
		}
	})

	apiGroup.GET("/ruleof72", func(ctx *gin.Context) {
		var cir compoundInterestRate
		if err := ctx.Bind(&cir); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(cir, ctx)
		}
	})

	apiGroup.GET("/simpleinterest", func(ctx *gin.Context) {
		var si simpleInterest
		if err := ctx.Bind(&si); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(si, ctx)
		}
	})

	apiGroup.GET("/variationofinvestment", func(ctx *gin.Context) {
		var vi variationOfInvestment
		if err := ctx.Bind(&vi); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			buildCalculationResponse(vi, ctx)
		}
	})
}
