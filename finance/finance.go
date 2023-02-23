package finance

func calculateNetWorth(netWorth netWorthInput) float64 {
	return netWorth.Assets - netWorth.Debts
}

func calculateCompoundInterest(compoundInterest compoundInterestInput) float64 {
	return compoundInterest.Principal *
		(1 + compoundInterest.AnnualInterestRate/
			compoundInterest.NumberTimesCompoundedPerYear) *
		(compoundInterest.NumberTimesCompoundedPerYear * compoundInterest.LengthBorrowedInYears)
}

func calculatePriceToEarningsRatio(peRatio peRatioInput) float64 {
	return peRatio.PricePerShare / peRatio.EarningsPerShare
}

func calculateBreakEvenPoint(breakEvenPoint breakEvenPointInput) float64 {
	return breakEvenPoint.FixedExpenses / breakEvenPoint.GrossProfitMarginInPercentage
}

func calculateNetIncome(netIncome netIncomeInput) float64 {
	return netIncome.Revenue - netIncome.Expenses
}
