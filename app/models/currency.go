package models

type Currency struct {
	BaseModel
	TenantID  string
	Name      string
	Symbol    string
	IsPrimary bool
	Code      string
	Rate      float64
}
