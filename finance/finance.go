package finance

import "math"

type calculation interface {
	calculate()
}

func (blr *basicLiquidityRatio) calculate() {
	blr.Result = blr.MonetaryAssets / blr.MonthlyExpenses
}

func (bep *breakEvenPoint) calculate() {
	bep.Result = bep.FixedExpenses / bep.GrossProfitMarginInPercentage
}

func (cf *cashFlow) calculate() {
	cf.Result = cf.Income - cf.Expenses
}

func (ci *compoundInterest) calculate() {
	ci.Result = ci.Principal *
		math.Pow(1+(ci.AnnualInterestRate/ci.NumberTimesCompoundedPerYear),
			ci.NumberTimesCompoundedPerYear*ci.LengthBorrowedInYears)
}

func (ci *ruleOf72) calculate() {
	ci.Result = 72 / ci.CompoundInterestRate
}

func (ni *netIncome) calculate() {
	ni.Result = ni.Revenue - ni.Expenses
}

func (nw *netWorth) calculate() {
	nw.Result = nw.Assets - nw.Debts
}

func (pe *peRatio) calculate() {
	pe.Result = pe.PricePerShare / pe.EarningsPerShare
}

func (si *simpleInterest) calculate() {
	si.Result = si.Principal *
		si.AnnualInterestRate *
		si.LengthBorrowedInYears
}

func (vi *variationOfInvestment) calculate() {
	vi.Result = (vi.CurrentPrice - vi.PurchasePrice) / vi.PurchasePrice
}
