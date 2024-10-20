package conf

type (
	UserIDKey        string
	BudgetAccountKey string
	BudgetKey        string
	IncomeKey        string
	ExpenseKey       string
)

const (
	UserIDContext        = UserIDKey("userID")
	BudgetAccountContext = BudgetAccountKey("budgetAccount")
	BudgetContext        = BudgetKey("budget")
	IncomeContext        = BudgetKey("income")
	ExpenseContext       = BudgetKey("expense")
)
