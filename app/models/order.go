package models

type Order struct {
	BaseModel
	TenantID      string
	CustomerID    string
	BranchID      string
	UserID        string
	OrderType     string // sale, return, exchange
	Subtotal      float64
	Total         float64
	PaymentStatus string
	PaymentMethod string
	Note          string
	Status        OrderStatus
	// relations
	// branche, user, customer, Tenant
}
type OrderStatus uint8

const (
	OrderStatusSuspend OrderStatus = iota
	OrderStatusPaid
	OrderStatusUnPaid
)

func (s OrderStatus) String() string {
	return [...]string{"suspend", "paid", "unpaid"}[s]
}
