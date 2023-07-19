package finance

import "math"

type Calculation interface {
	calculate()
}

func (blr *BasicLiquidityRatio) calculate() {
	blr.Result = blr.MonetaryAssets / blr.MonthlyExpenses
}

func (bep *BreakEvenPoint) calculate() {
	bep.Result = bep.FixedExpenses / bep.GrossProfitMarginInPercentage
}

func (cf *CashFlow) calculate() {
	cf.Result = cf.Income - cf.Expenses
}

func (ci *CompoundInterest) calculate() {
	ci.Result = ci.Principal *
		math.Pow(1+(ci.AnnualInterestRate/ci.NumberTimesCompoundedPerYear),
			ci.NumberTimesCompoundedPerYear*ci.LengthBorrowedInYears)
}

func (ci *RuleOf72) calculate() {
	ci.Result = 72 / ci.CompoundInterestRate
}

func (ni *NetIncome) calculate() {
	ni.Result = ni.Revenue - ni.Expenses
}

func (nw *NetWorth) calculate() {
	nw.Result = nw.Assets - nw.Debts
}

func (pe *PERatio) calculate() {
	pe.Result = pe.PricePerShare / pe.EarningsPerShare
}

func (si *SimpleInterest) calculate() {
	si.Result = si.Principal *
		si.AnnualInterestRate *
		si.LengthBorrowedInYears
}

func (vi *VariationOfInvestment) calculate() {
	vi.Result = (vi.CurrentPrice - vi.PurchasePrice) / vi.PurchasePrice
}
