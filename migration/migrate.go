package migrate

import (
	"WMS/models"
	"fmt"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	fmt.Println("starting migrations!")
	err := db.AutoMigrate(&models.Hub{}, &models.Inventory{}, &models.SKU{}, &models.Product{})
	if err != nil {
		fmt.Println("migration unsuccessful!", err)
	}
	fmt.Println("migrations completed successfully")
}
