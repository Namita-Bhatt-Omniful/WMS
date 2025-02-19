package controllers

import (
	"WMS/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func ViewSKU(ctx *gin.Context) {
// 	id, err1 := strconv.Atoi(ctx.Param("id"))
// 	if err1 != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong parameters passed"})
// 		return
// 	}
// 	sku, err := repositories.FindSKU(id) // find SKU corresponding to this id
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, sku) // return sku details as json
// }

type Order struct {
	OrderID       string `json:"id"`
	SellerID      string `json:"seller_id"`
	SKUID         string `json:"sku_id"`
	ItemCount     string `json:"item_count"`
	ModeOfPayment string `json:"mode_of_payment"`
	Status        string `json:"status"`
}

type SKU struct {
	ID string `gorm:"primaryKey"`
}
type response struct {
	ValidOrders []Order `json:"valid_orders"`
}

func VerifySKU(ctx *gin.Context) {

	var orders []Order
	// bind the incoming JSON request body to the orders variable
	if err := ctx.ShouldBindJSON(&orders); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// fmt.Printf("%+v", orders)

	db := config.DB
	var validOrders []Order
	for _, order := range orders {
		var sku SKU
		if err := db.First(&sku, "id = ?", order.SKUID).Error; err != nil {
			fmt.Println("Error finding sku:", err)
		}
		validOrders = append(validOrders, order)
	}
	// fmt.Println(validOrders)
	resp := &response{
		ValidOrders: validOrders,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"is_success":  true,
		"status_code": 200,
		"data":        resp,
		"meta":        gin.H{},
	})
}
