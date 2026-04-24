package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Account struct {
	orm.Model
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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
