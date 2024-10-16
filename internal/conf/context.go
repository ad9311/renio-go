package conf

type UserIDKey string
type BudgetAccountKey string
type BudgetKey string
type IncomeKey string

const UserIDContext = UserIDKey("userID")
const BudgetAccountContext = BudgetAccountKey("budgetAccount")
const BudgetContext = BudgetKey("budget")
const IncomeContext = BudgetKey("income")
