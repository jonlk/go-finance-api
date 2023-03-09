package finance

type BasicLiquidityRatio struct {
	MonetaryAssets  float64 `form:"monetaryAssets" json:"monetaryAssets" binding:"required"`
	MonthlyExpenses float64 `form:"monthlyExpenses" json:"monthlyExpenses" binding:"required"`
	Result          float64 `json:"result"`
}

type BreakEvenPoint struct {
	FixedExpenses                 float64 `form:"fixedExpenses" json:"fixedExpenses" binding:"required"`
	GrossProfitMarginInPercentage float64 `form:"grossProfitMarginInPercentage" json:"grossProfitMarginInPercentage" binding:"required"`
	Result                        float64 `json:"result"`
}

type CashFlow struct {
	Income   float64 `form:"income" json:"income" binding:"required"`
	Expenses float64 `form:"expenses" json:"expenses" binding:"required"`
	Result   float64 `json:"result"`
}

type CompoundInterest struct {
	Principal                    float64 `form:"principal" json:"principal" binding:"required"`
	AnnualInterestRate           float64 `form:"annualInterestRate" json:"annualInterestRate" binding:"required"`
	NumberTimesCompoundedPerYear float64 `form:"numberTimesCompoundedPerYear" json:"numberTimesCompoundedPerYear" binding:"required"`
	LengthBorrowedInYears        float64 `form:"lengthBorrowedInYears" json:"lengthBorrowedInYears" binding:"required"`
	Result                       float64 `json:"result"`
}

type RuleOf72 struct {
	CompoundInterestRate float64 `form:"compoundInterestRate" json:"compoundInterestRate" binding:"required"`
	Result               float64 `json:"result"`
}

type NetIncome struct {
	Revenue  float64 `form:"revenue" json:"revenue" binding:"required"`
	Expenses float64 `form:"expenses" json:"expenses" binding:"required"`
	Result   float64 `json:"result"`
}

type NetWorth struct {
	Assets float64 `form:"assets" json:"assets" binding:"required"`
	Debts  float64 `form:"debts" json:"debts" binding:"required"`
	Result float64 `json:"result"`
}

type PERatio struct {
	PricePerShare    float64 `form:"pricePerShare" json:"pricePerShare" binding:"required"`
	EarningsPerShare float64 `form:"earningsPerShare" json:"earningsPerShare" binding:"required"`
	Result           float64 `json:"result"`
}

type SimpleInterest struct {
	Principal             float64 `form:"principal" json:"principal" binding:"required"`
	AnnualInterestRate    float64 `form:"annualInterestRate" json:"annualInterestRate" binding:"required"`
	LengthBorrowedInYears float64 `form:"lengthBorrowedInYears" json:"lengthBorrowedInYears" binding:"required"`
	Result                float64 `json:"result"`
}

type VariationOfInvestment struct {
	CurrentPrice  float64 `form:"currentPrice" json:"currentPrice" binding:"required"`
	PurchasePrice float64 `form:"purchasePrice" json:"purchasePrice" binding:"required"`
	Result        float64 `json:"result"`
}
