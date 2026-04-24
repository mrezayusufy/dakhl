package models

type Warehouse struct {
	ID
	Name     string
	TenantID string
	BranchID string
}

// Product has many WareHouse
// Warehouse has many products
// Many To Many
