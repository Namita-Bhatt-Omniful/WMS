package repositories

import (
	"WMS/config"
	"WMS/models"
)

func FindHub(id string) (*models.Hub, error) {
	var hub models.Hub
	if err := config.DB.First(&hub, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &hub, nil
}
