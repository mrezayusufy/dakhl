package models

type Customer struct {
	BaseModel
	TenantID    string
	BranchID    string
	Name        string
	Email       string
	Phone       string
	Address     string
	CreditLimit float64
	Note        string
	IsActive    bool
	IsPublic    bool
}
