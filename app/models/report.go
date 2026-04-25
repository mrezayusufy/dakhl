package models

import (
	"encoding/json"
	"time"
)

type Report struct {
	BaseModel
	TenantID   string
	UserID     string
	ReportType string // sales, invetory, purchases, financial
	Title      string
	Parameters json.RawMessage
	FileUrl    string
	CreatedAt  time.Time
}
