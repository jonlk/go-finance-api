package finance

import (
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type calculationResponse struct {
	ResponseId      uuid.UUID   `json:"responseId"`
	Timestamp       int64       `json:"timestamp"`
	CalculationType string      `json:"calculationType"`
	Result          calculation `json:"result"`
}

func BuildCalculationResponse(c calculation, ctx *gin.Context) {
	c.calculate()

	calcType := reflect.
		PointerTo(reflect.TypeOf(c)).
		String()

	response := calculationResponse{
		ResponseId:      uuid.New(),
		Timestamp:       time.Now().Unix(),
		CalculationType: strings.Split(calcType, ".")[1],
		Result:          c,
	}
	ctx.JSON(http.StatusOK, &response)
}
