package models

import (
	"github.com/goravel/framework/database/orm"
)

type Unit struct {
	orm.Model
	TenantID uint
	Name     string
	Symbol   string
	Value    uint
}
