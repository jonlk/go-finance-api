package finance

func calculateNetWorth(nw netWorth) float64 {
	result := nw.Assets - nw.Debts
	return result
}

func calculateCompoundInterest(ci compoundInterest) float64 {
	result := ci.Principal *
		(1 + (ci.AnnualInterestRate / ci.NumberTimesCompoundedPerYear)) *
		(ci.NumberTimesCompoundedPerYear * ci.LengthBorrowedInYears)
	return result
}

func calculatePriceToEarningsRatio(pe peRatio) float64 {
	result := pe.PricePerShare / pe.EarningsPerShare
	return result
}

func calculateBreakEvenPoint(bep breakEvenPoint) float64 {
	result := bep.FixedExpenses / bep.GrossProfitMarginInPercentage
	return result
}

func calculateNetIncome(ni netIncome) float64 {
	result := ni.Revenue - ni.Expenses
	return result
}

func calculateCashFlow(cf cashFlow) float64 {
	result := cf.Income - cf.Expenses
	return result
}

func calculateSimpleInterest(si simpleInterest) float64 {
	result := si.Principal *
		si.AnnualInterestRate *
		si.LengthBorrowedInYears
	return result
}

func calculateVariationOfInvestment(vi variationOfInvestment) float64 {
	result := (vi.CurrentPrice - vi.PurchasePrice) / vi.PurchasePrice
	return result
}
