package models

type ProductPhoto struct {
	ID
	ProductID uint
	Uri       string
	IsPrimary bool
	SortOrder uint
}
