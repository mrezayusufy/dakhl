package models

type OrderItem struct {
	ID
	OrderID   string
	ProductID string
	Qty       uint
	UnitPrice float64
	Total     float64
}
