package vars

const (
	UserIDKey        = ContextKey("userID")
	BudgetAccountKey = ContextKey("budgetAccount")
	BudgetKey        = ContextKey("budget")
	IncomeKey        = ContextKey("income")
	ExpenseKey       = ContextKey("expense")
)

type ContextKey string

var FilteredFields = map[string]bool{
	"Password":             true,
	"PasswordConfirmation": true,
}
