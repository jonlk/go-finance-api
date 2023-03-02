package finance

type basicLiquidityRatio struct {
	MonetaryAssets  float64 `form:"monetaryAssets" json:"monetaryAssets" binding:"required"`
	MonthlyExpenses float64 `form:"monthlyExpenses" json:"monthlyExpenses" binding:"required"`
	Result          float64 `json:"result"`
}

type breakEvenPoint struct {
	FixedExpenses                 float64 `form:"fixedExpenses" json:"fixedExpenses" binding:"required"`
	GrossProfitMarginInPercentage float64 `form:"grossProfitMarginInPercentage" json:"grossProfitMarginInPercentage" binding:"required"`
	Result                        float64 `json:"result"`
}

type cashFlow struct {
	Income   float64 `form:"income" json:"income" binding:"required"`
	Expenses float64 `form:"expenses" json:"expenses" binding:"required"`
	Result   float64 `json:"result"`
}

type compoundInterest struct {
	Principal                    float64 `form:"principal" json:"principal" binding:"required"`
	AnnualInterestRate           float64 `form:"annualInterestRate" json:"annualInterestRate" binding:"required"`
	NumberTimesCompoundedPerYear float64 `form:"numberTimesCompoundedPerYear" json:"numberTimesCompoundedPerYear" binding:"required"`
	LengthBorrowedInYears        float64 `form:"lengthBorrowedInYears" json:"lengthBorrowedInYears" binding:"required"`
	Result                       float64 `json:"result"`
}

type ruleOf72 struct {
	CompoundInterestRate float64 `form:"compoundInterestRate" json:"compoundInterestRate" binding:"required"`
	Result               float64 `json:"result"`
}

type netIncome struct {
	Revenue  float64 `form:"revenue" json:"revenue" binding:"required"`
	Expenses float64 `form:"expenses" json:"expenses" binding:"required"`
	Result   float64 `json:"result"`
}

type netWorth struct {
	Assets float64 `form:"assets" json:"assets" binding:"required"`
	Debts  float64 `form:"debts" json:"debts" binding:"required"`
	Result float64 `json:"result"`
}

type peRatio struct {
	PricePerShare    float64 `form:"pricePerShare" json:"pricePerShare" binding:"required"`
	EarningsPerShare float64 `form:"earningsPerShare" json:"earningsPerShare" binding:"required"`
	Result           float64 `json:"result"`
}

type simpleInterest struct {
	Principal             float64 `form:"principal" json:"principal" binding:"required"`
	AnnualInterestRate    float64 `form:"annualInterestRate" json:"annualInterestRate" binding:"required"`
	LengthBorrowedInYears float64 `form:"lengthBorrowedInYears" json:"lengthBorrowedInYears" binding:"required"`
	Result                float64 `json:"result"`
}

type variationOfInvestment struct {
	CurrentPrice  float64 `form:"currentPrice" json:"currentPrice" binding:"required"`
	PurchasePrice float64 `form:"purchasePrice" json:"purchasePrice" binding:"required"`
	Result        float64 `json:"result"`
}
