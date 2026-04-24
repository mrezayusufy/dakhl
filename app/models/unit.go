package models

type Unit struct {
	ID
	TenantID uint
	Name     string
	Symbol   string
	Value    uint
}
