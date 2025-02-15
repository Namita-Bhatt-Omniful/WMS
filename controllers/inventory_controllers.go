package controllers

import (
	"WMS/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	ctx.Param("update")
}
