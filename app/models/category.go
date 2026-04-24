package models

import (
	"github.com/goravel/framework/database/orm"
)

type Category struct {
	orm.Model
	TenantID    string
	ParentID    string
	Name        string
	Color       string
	Photo       string
	Description string
}
