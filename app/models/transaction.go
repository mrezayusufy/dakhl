package models

type Transaction struct {
	BaseModel
	TenantID        string
	BranchID        string
	UserID          string
	TransactionType string
	CategoryTrxID   string
	Amount          float64
	CurrencyID      string
	Status          string // pending, completed, failed, reversed, cancalled
	PaymentMethod   string // cash, card, check, bank_transfer, wallet
	ReferenceType   string // order, purchase, expence, invoice, etc,
	ReferenceID     string
	PartyType       string // customer, vendor, user, bank, cash
	PartyID         string
	Note            string
	ReferenceNumber string
	IsIcome         bool
	AccountID       string
}
