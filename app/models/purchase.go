package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Purchase struct {
	orm.Model
	TenandID       uint
	BranchID       uint
	VendorID       uint
	UserID         uint
	PurchaseNumber uint
	PurchaseType   string // purchase - return
	Subtotal       float64
	ExtraAmount    float64
	Total          float64
	PaymentStatus  string // pending, partial, paid, cancelled
	PaymentMethod  string
	Status         string // pending, confirmed, received, returned, cancelled
	DueDate        time.Time
	Note           string
	ReceivedAt     time.Time
}
