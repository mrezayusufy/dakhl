package models

type Notification struct {
	BaseModel
	TenantID string
	BranchID string
	UserID   string
	Type     string
	Title    string
	Message  string
	IsRead   bool
}
