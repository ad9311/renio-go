package vars

const (
	CurrentUserKey   = ContextKey("currentUser")
	UserSignedInKey  = ContextKey("userSignedIn")
	UserIDKey        = ContextKey("userIDKey")
	BudgetAccountKey = ContextKey("budgetAccount")
	BudgetKey        = ContextKey("budget")
	IncomeKey        = ContextKey("income")
	ExpenseKey       = ContextKey("expense")
	AppDataKey       = ContextKey("appData")

	EmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

type ContextKey string
