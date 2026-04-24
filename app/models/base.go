package models

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type ID struct {
	ID string `json:"id" gorm:"type:varchar(12);primaryKey"`
}
type BaseModel struct {
	ID
	orm.Timestamps
}

// BeforeCreate hook برای تولید خودکار ID
func (b *ID) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = GenerateShortID()
	}
	return nil
}

// GenerateShortID تولید شناسه یکتا
func GenerateShortID() string {
	bytes := make([]byte, 9)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)[:12]
}
