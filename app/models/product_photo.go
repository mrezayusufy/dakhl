package models

import (
	"github.com/goravel/framework/database/orm"
)

type ProductPhoto struct {
	orm.Model
	ProductID uint
	Uri       string
	IsPrimary bool
	SortOrder uint
}
