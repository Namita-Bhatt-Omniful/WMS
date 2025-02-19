package models

import "time"

type Inventory struct {
	ID        int       `gorm:"type:integer;primaryKey;not null" json:"id"`
	HubID     int       `gorm:"type:integer;not null" json:"hub_id"`
	SKUID     int       `gorm:"type:integer;not null" json:"sku_id"`
	SellerID  int       `gorm:"type:integer;not null" json:"seller_id"`
	Quantity  int       `gorm:"type:integer" json:"qty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
