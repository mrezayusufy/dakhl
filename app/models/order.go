package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Order struct {
	orm.Model
	Name    string
	license string
	Logo    string
	Note    string
	Address string
	Status  OrderStatus
	// relations
	// branches, users, products, customers, vendors, categories, trx,
	CreatedAt time.Time
	UpdatedAt time.Time
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
