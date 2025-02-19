package controllers

import (
	"WMS/config"
	"WMS/models"
	"WMS/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	SKUID    int `json:"sku_id"`
	Quantity int `json:"qty"`
}
type resp struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func GetInventory(ctx *gin.Context) {
	// returns error in case of empty string conversion
	sellerID, err1 := strconv.Atoi(ctx.Query("seller_id"))
	hubID, err2 := strconv.Atoi(ctx.Query("hub_id"))
	if err1 != nil || err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong parameters passed"})
		return
	}
	inventory, err := repositories.FindInventory(sellerID, hubID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There is some error while fetching inventory details"})
		return
	}

	ctx.JSON(http.StatusOK, inventory)
}

func UpdateInventory(ctx *gin.Context) {
	var req request
	// bind the incoming JSON request body to the req variable
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Printf("%+v", req)
	var inventory models.Inventory
	db := config.DB
	err := db.Where("sku_id = ?", req.SKUID).First(&inventory).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	if inventory.Quantity < req.Quantity {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stock is currently unavailable"})
		return
	}
	inventory.Quantity -= req.Quantity
	if err := db.Save(&inventory).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update inventory"})
		return
	}
	res := &resp{
		Message: "inventory updated successfully",
		Status:  "http.StatusOK",
	}
	ctx.JSON(http.StatusOK, gin.H{
		"is_success":  true,
		"status_code": 200,
		"data":        res,
		"meta":        gin.H{},
	})
}
