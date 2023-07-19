package finance

import (
	"reflect"
	"strings"
)

type FinanceResult struct {
	CalculationType string      `json:"calculationType"`
	Data            Calculation `json:"data"`
}

func GetFinanceResult(calculation Calculation) FinanceResult {
	calculation.calculate()

	calcType := reflect.
		PointerTo(reflect.TypeOf(calculation)).
		String()

	return FinanceResult{
		CalculationType: strings.Split(calcType, ".")[1],
		Data:            calculation,
	}
}
