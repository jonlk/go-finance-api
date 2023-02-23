package finance

func calculateNetWorth(nw netWorth) float64 {
	return nw.Assets - nw.Debts
}

func calculateCompoundInterest(ci compoundInterest) float64 {
	return ci.Principal *
		(1 + (ci.AnnualInterestRate / ci.NumberTimesCompoundedPerYear)) *
		(ci.NumberTimesCompoundedPerYear * ci.LengthBorrowedInYears)
}

func calculatePriceToEarningsRatio(pe peRatio) float64 {
	return pe.PricePerShare / pe.EarningsPerShare
}

func calculateBreakEvenPoint(bep breakEvenPoint) float64 {
	return bep.FixedExpenses / bep.GrossProfitMarginInPercentage
}

func calculateNetIncome(ni netIncome) float64 {
	return ni.Revenue - ni.Expenses
}

func calculateCashFlow(cf cashFlow) float64 {
	return cf.Income - cf.Expenses
}

func calculateSimpleInterest(si simpleInterest) float64 {
	return si.Principal *
		si.AnnualInterestRate *
		si.LengthBorrowedInYears
}

func calculateVariationOfInvestment(vi variationOfInvestment) float64 {
	return (vi.CurrentPrice - vi.PurchasePrice) / vi.PurchasePrice
}
