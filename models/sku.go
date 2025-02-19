package models

import "time"

type SKU struct {
	ID          int       `gorm:"type:integer;primaryKey;not null" json:"id"`
	ProductID   int       `gorm:"type:integer;not null" json:"hub_id"`
	SkuCodes    int       `gorm:"type:integer" json:"sku_codes"`
	Name        string    `gorm:"type:varchar(50)" json:"name"`
	UnitPrice   int       `gorm:"type:integer" json:"price"`
	Fragile     bool      `gorm:"type:boolean" json:"fragile"`
	Description string    `gorm:"type:varchar(1000)" json:"description"`
	BoxSizing   string    `gorm:"type:varchar(20)" json:"box_size"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
