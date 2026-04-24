package models

type Category struct {
	ID
	TenantID    string
	ParentID    string
	Name        string
	Color       string
	Photo       string
	Description string
}
