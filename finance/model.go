package finance

type basicLiquidityRatio struct {
	MonetaryAssets  float64 `form:"monetaryAssets" binding:"required"`
	MonthlyExpenses float64 `form:"monthlyExpenses" binding:"required"`
}

type breakEvenPoint struct {
	FixedExpenses                 float64 `form:"fixedExpenses" binding:"required"`
	GrossProfitMarginInPercentage float64 `form:"grossProfitMarginInPercentage" binding:"required"`
}

type cashFlow struct {
	Income   float64 `form:"income" binding:"required"`
	Expenses float64 `form:"expenses" binding:"required"`
}

type compoundInterest struct {
	Principal                    float64 `form:"principal" binding:"required"`
	AnnualInterestRate           float64 `form:"annualInterestRate" binding:"required"`
	NumberTimesCompoundedPerYear float64 `form:"numberTimesCompoundedPerYear" binding:"required"`
	LengthBorrowedInYears        float64 `form:"lengthBorrowedInYears" binding:"required"`
}

type compoundInterestRate float64

type netIncome struct {
	Revenue  float64 `form:"revenue" binding:"required"`
	Expenses float64 `form:"expenses" binding:"required"`
}

type netWorth struct {
	Assets float64 `form:"assets" binding:"required"`
	Debts  float64 `form:"debts" binding:"required"`
}

type peRatio struct {
	PricePerShare    float64 `form:"pricePerShare" binding:"required"`
	EarningsPerShare float64 `form:"earningsPerShare" binding:"required"`
}

type simpleInterest struct {
	Principal             float64 `form:"principal" binding:"required"`
	AnnualInterestRate    float64 `form:"annualInterestRate" binding:"required"`
	LengthBorrowedInYears float64 `form:"lengthBorrowedInYears" binding:"required"`
}

type variationOfInvestment struct {
	CurrentPrice  float64 `form:"currentPrice" binding:"required"`
	PurchasePrice float64 `form:"purchasePrice" binding:"required"`
}
