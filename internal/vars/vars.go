package vars

const (
	AllowedJWTKey    = ContextKey("allowedJWTID")
	UserIDKey        = ContextKey("userID")
	BudgetAccountKey = ContextKey("budgetAccount")
	BudgetKey        = ContextKey("budget")
	IncomeKey        = ContextKey("income")
	ExpenseKey       = ContextKey("expense")

	EmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

type ContextKey string

var FilteredFields = map[string]bool{
	"Password":             true,
	"PasswordConfirmation": true,
}
