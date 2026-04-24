package models

import (
	"encoding/json"

	"github.com/goravel/framework/database/orm"
)

type Tanent struct {
	orm.Model
	Name        string
	license     string
	Logo        string
	Description string
	Address     string
	Phone       string
	Settings    json.RawMessage
	Status      TenantStatus
	IsActive    bool
	// relations
	// branches, users, products, customers, vendors, categories, trx,
}
type TenantStatus uint8

const (
	TenantStatusSuspend TenantStatus = iota
	TenantStatusActive
	TenantStatusTrial
)

func (s TenantStatus) String() string {
	return [...]string{"suspend", "active", "trial"}[s]
}
