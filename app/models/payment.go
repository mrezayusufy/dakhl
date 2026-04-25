package models

type Payment struct {
	BaseModel
	TenantID    string
	ParentID    string
	Name        string
	AccountType string // asset, liability, equity, revenue, expense
	IsCash      bool
	IsBank      bool
	Balance     float64
	IsActive    bool
}
