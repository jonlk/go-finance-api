package finance

type netWorthInput struct {
	Assets float64 `json:"assets"`
	Debts  float64 `json:"debts"`
}

type compoundInterestInput struct {
	Principal                    float64 `json:"principal"`
	AnnualInterestRate           float64 `json:"annualInterestRate"`
	NumberTimesCompoundedPerYear float64 `json:"numberTimesCompoundedPerYear"`
	LengthBorrowedInYears        float64 `json:"lengthBorrowedInYears"`
}

type peRatioInput struct {
	PricePerShare    float64 `json:"pricePerShare"`
	EarningsPerShare float64 `json:"earningsPerShare"`
}

type breakEvenPointInput struct {
	FixedExpenses                 float64 `json:"fixedExpenses"`
	GrossProfitMarginInPercentage float64 `json:"grossProfitMarginInPercentage"`
}

type netIncomeInput struct {
	Revenue  float64 `json:"revenue"`
	Expenses float64 `json:"expenses"`
}
