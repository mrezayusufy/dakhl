package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/goravel/framework/database/orm"
)

type User struct {
	BaseModel
	Name           string
	TenantID       string `json:"tenant_id"`
	BranchID       string `json:"branch_id"`
	Email          string
	Phone          string
	HashedPassword string `json:"hashed_password"`
	Avatar         string
	IsActive       bool
	LastLogin      time.Time
	Role           string // admin, manager, cashier
	Permissions    json.RawMessage
	Tags           []UserTag `gorm:"serializer:json"`
	orm.SoftDeletes
}

type UserTag struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

func (r *UserTag) Scan(value any) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, r)
}

func (r *UserTag) Value() (driver.Value, error) {
	return json.Marshal(r)
}
