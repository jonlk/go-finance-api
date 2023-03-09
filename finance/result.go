package finance

import (
	"reflect"
	"strings"
)

type FinanceResult struct {
	CalculationType string      `json:"calculationType"`
	Data            calculation `json:"data"`
}

func GetFinanceResult(c calculation) FinanceResult {
	c.calculate()

	calcType := reflect.
		PointerTo(reflect.TypeOf(c)).
		String()

	return FinanceResult{
		CalculationType: strings.Split(calcType, ".")[1],
		Data:            c,
	}
}
