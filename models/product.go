package models

import "time"

type Product struct {
	ID           int       `gorm:"type:integer;primaryKey;not null" json:"id"`
	SellerID     int       `gorm:"type:integer;not null" json:"seller_id"`
	Name         string    `gorm:"type:varchar(50)" json:"name"`
	Brand        string    `gorm:"type:varchar(50)" json:"brand"`
	Manufacturer string    `gorm:"type:varchar(50)" json:"manufacturer"`
	Description  string    `gorm:"type:varchar(1000)" json:"description"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
