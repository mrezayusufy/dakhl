package models

import "time"

type Expense struct {
	BaseModel
	TenantID      string
	BranchID      string
	UserID        string
	ExpenseType   string // rent, salary, utility, marketing, etc
	Title         string
	Amount        float64
	PaymentMethod string // card, cash, debit
	Status        string // pending, approved, paid, cancelled
	ExpenseDate   time.Time
	DueDate       time.Time
	PaidAt        time.Time
	Note          string
}
