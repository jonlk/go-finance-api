package finance

type basicLiquidityRatio struct {
	MonetaryAssets  float64 `form:"monetaryAssets"`
	MonthlyExpenses float64 `form:"monthlyExpenses"`
}

type breakEvenPoint struct {
	FixedExpenses                 float64 `form:"fixedExpenses"`
	GrossProfitMarginInPercentage float64 `form:"grossProfitMarginInPercentage"`
}

type cashFlow struct {
	Income   float64 `form:"income"`
	Expenses float64 `form:"expenses"`
}

type compoundInterest struct {
	Principal                    float64 `form:"principal"`
	AnnualInterestRate           float64 `form:"annualInterestRate"`
	NumberTimesCompoundedPerYear float64 `form:"numberTimesCompoundedPerYear"`
	LengthBorrowedInYears        float64 `form:"lengthBorrowedInYears"`
}

type compoundInterestRate float64

type netIncome struct {
	Revenue  float64 `form:"revenue"`
	Expenses float64 `form:"expenses"`
}

type netWorth struct {
	Assets float64 `form:"assets"`
	Debts  float64 `form:"debts"`
}

type peRatio struct {
	PricePerShare    float64 `form:"pricePerShare"`
	EarningsPerShare float64 `form:"earningsPerShare"`
}

type simpleInterest struct {
	Principal             float64 `form:"principal"`
	AnnualInterestRate    float64 `form:"annualInterestRate"`
	LengthBorrowedInYears float64 `form:"lengthBorrowedInYears"`
}

type variationOfInvestment struct {
	CurrentPrice  float64 `form:"currentPrice"`
	PurchasePrice float64 `form:"purchasePrice"`
}
