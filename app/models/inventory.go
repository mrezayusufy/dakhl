package models

type Inventory struct {
	BaseModel
	TenantID  string
	BranchID  string
	ProductID string
	Qty       float64
	MinStock  int
	MaxStock  int
}
