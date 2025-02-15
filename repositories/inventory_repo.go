package repositories

import (
	"WMS/config"
	"WMS/models"
	"fmt"
)

func FindInventory(sellerID, hubID int) (*[]models.Inventory, error) {
	var inventories []models.Inventory
	err := config.DB.Where("seller_id = ? AND hub_id= ?", sellerID, hubID).Find(&inventories).Error
	if err != nil {
		return nil, err
	}
	if len(inventories) == 0 {
		return nil, fmt.Errorf("No inventory found")
	}
	return &inventories, nil
}
