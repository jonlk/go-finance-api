package finance

import "math"

type calculation interface {
	calculate() float64
}

func (blr basicLiquidityRatio) calculate() float64 {
	result := blr.MonetaryAssets / blr.MonthlyExpenses
	return result
}

func (bep breakEvenPoint) calculate() float64 {
	result := bep.FixedExpenses / bep.GrossProfitMarginInPercentage
	return result
}

func (cf cashFlow) calculate() float64 {
	result := cf.Income - cf.Expenses
	return result
}

func (ci compoundInterest) calculate() float64 {
	result := ci.Principal *
		math.Pow(1+(ci.AnnualInterestRate/ci.NumberTimesCompoundedPerYear),
			ci.NumberTimesCompoundedPerYear*ci.LengthBorrowedInYears)
	return result
}

func (ci compoundInterestRate) calculate() float64 {
	result := 72 / ci
	return float64(result)
}

func (ni netIncome) calculate() float64 {
	result := ni.Revenue - ni.Expenses
	return result
}

func (nw netWorth) calculate() float64 {
	result := nw.Assets - nw.Debts
	return result
}

func (pe peRatio) calculate() float64 {
	result := pe.PricePerShare / pe.EarningsPerShare
	return result
}

func (si simpleInterest) calculate() float64 {
	result := si.Principal *
		si.AnnualInterestRate *
		si.LengthBorrowedInYears
	return result
}

func (vi variationOfInvestment) calculate() float64 {
	result := (vi.CurrentPrice - vi.PurchasePrice) / vi.PurchasePrice
	return result
}
