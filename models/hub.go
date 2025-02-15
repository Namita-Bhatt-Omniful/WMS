package models

import (
	"time"
)

type Hub struct {
	ID        int       `gorm:"type:integer;primaryKey;not null" json:"id"`
	TenantID  int       `gorm:"type:integer;not null" json:"tenant_id"`
	Name      string    `gorm:"type:varchar(50)" json:"name"`
	Address   string    `gorm:"type:varchar(100)" json:"address"`
	City      string    `gorm:"type:varchar(50)" json:"city"`
	State     string    `gorm:"type:varchar(50)" json:"state"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
