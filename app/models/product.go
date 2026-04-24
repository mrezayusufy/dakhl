package models

import (
	"time"
)

type Product struct {
	BaseModel
	Name         string
	Title        *string
	Description  *string
	Barcode      *string
	SKU          *string
	MinStock     *uint
	Category     *uint
	Units        []*Unit `gorm:"one2many:product_units;"`
	Image        string
	SalePrice    *float64
	BuyPrice     *float64
	CostPrice    *float64
	Qty          uint
	DepartmentID string
	WareHouseID  string
	VendorID     string
	BranchID     uint
	IsFav        bool
	Expiry       time.Time
}
