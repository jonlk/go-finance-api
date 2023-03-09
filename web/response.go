package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jonlk/go-finance-api/finance"
)

type responseModel struct {
	ResponseId uuid.UUID              `json:"responseId"`
	Timestamp  int64                  `json:"timestamp"`
	Result     *finance.FinanceResult `json:"result"`
}

func buildResponseModel(result *finance.FinanceResult, ctx *gin.Context) {

	response := &responseModel{
		ResponseId: uuid.New(),
		Timestamp:  time.Now().Unix(),
		Result:     result,
	}

	ctx.JSON(http.StatusOK, response)
}
