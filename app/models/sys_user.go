package models

import (
	"encoding/json"
)

type SysUser struct {
	BaseModel
	Name            string
	Email           string
	EncryptPassword string
	Role            string
	Permissions     json.RawMessage
}
