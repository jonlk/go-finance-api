package finance

type basicLiquidityRatio struct {
	MonetaryAssets  float64 `json:"monetaryAssets"`
	MonthlyExpenses float64 `json:"monthlyExpenses"`
}

type breakEvenPoint struct {
	FixedExpenses                 float64 `json:"fixedExpenses"`
	GrossProfitMarginInPercentage float64 `json:"grossProfitMarginInPercentage"`
}

type cashFlow struct {
	Income   float64 `json:"income"`
	Expenses float64 `json:"expenses"`
}

type compoundInterest struct {
	Principal                    float64 `json:"principal"`
	AnnualInterestRate           float64 `json:"annualInterestRate"`
	NumberTimesCompoundedPerYear float64 `json:"numberTimesCompoundedPerYear"`
	LengthBorrowedInYears        float64 `json:"lengthBorrowedInYears"`
}

type compoundInterestRate float64

type netIncome struct {
	Revenue  float64 `json:"revenue"`
	Expenses float64 `json:"expenses"`
}

type netWorth struct {
	Assets float64 `json:"assets"`
	Debts  float64 `json:"debts"`
}

type peRatio struct {
	PricePerShare    float64 `json:"pricePerShare"`
	EarningsPerShare float64 `json:"earningsPerShare"`
}

type simpleInterest struct {
	Principal             float64 `json:"principal"`
	AnnualInterestRate    float64 `json:"annualInterestRate"`
	LengthBorrowedInYears float64 `json:"lengthBorrowedInYears"`
}

type variationOfInvestment struct {
	CurrentPrice  float64 `json:"currentPrice"`
	PurchasePrice float64 `json:"purchasePrice"`
}
