package models

import (
	"time"
)

type PurchaseItem struct {
	ID
	PurchaseID  uint
	ProductID   uint
	Qty         uint
	ReceivedQty uint
	UnitPrice   uint
	Total       float64
	ExpiryDate  time.Time
}
