package models

type Account struct {
	BaseModel
	TenantID      string
	ParentID      string
	Name          string
	AccountType   string
	IsCash        bool
	IsBank        bool
	BankName      string
	AccountNumber string
	Balance       float64
	IsActive      bool
}
