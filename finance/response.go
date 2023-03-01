package finance

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
