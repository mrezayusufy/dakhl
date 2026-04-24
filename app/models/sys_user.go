package models

import (
	"encoding/json"

	"github.com/goravel/framework/database/orm"
)

type SysUser struct {
	orm.Model
	Name            string
	Email           string
	EncryptPassword string
	Role            string
	Permissions     json.RawMessage
}
