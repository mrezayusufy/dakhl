package models

type InventoryMovement struct {
	BaseModel
	TenantID      string
	UserID        string
	MovementType  string // purchase, sale, return, adjustment, transfer, damage
	ReferenceType string // order, purchase, damage_report
	ReferenceID   string
	BranchID      string
	ToBranchID    string
	ProductID     string
	Qty           int // input = positive, output = negative
	QtyBefore     int
	QtyAfter      int
	UnitCost      float64
	TotalCost     float64
	Note          string
}
