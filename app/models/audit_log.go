package models

import "encoding/json"

type AuditLog struct {
	BaseModel
	TenantID   string
	UserID     string
	Action     string // create, update, delete, login, etc.
	EntityType string
	EntityID   string
	OldValues  json.RawMessage
	NewValues  json.RawMessage
	IP         string
}
