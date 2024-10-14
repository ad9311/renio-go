package conf

type UserIDKey string
type BudgetKey string
type BudgetAccountKey string

const UserIDContext = UserIDKey("userID")
const BudgetContext = BudgetKey("budget")
const BudgetAccountContext = BudgetAccountKey("budgetAccount")
