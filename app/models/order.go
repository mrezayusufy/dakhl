package models

type Order struct {
	BaseModel
	Name    string
	license string
	Logo    string
	Note    string
	Address string
	Status  OrderStatus
	// relations
	// branches, users, products, customers, vendors, categories, trx,
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
