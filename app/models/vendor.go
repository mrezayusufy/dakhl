package models

type Vendor struct {
	BaseModel
	TenantID string
	BranchID string
	Name     string
	Phone    string
	Address  string
	Avatar   string
	Note     string
	IsActive bool
}
