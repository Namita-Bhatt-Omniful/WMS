package controllers

import (
	"WMS/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHubByID(ctx *gin.Context) {
	id := ctx.Param("id")
	hub, err := repositories.FindHub(id) // find hub corresponding to this id
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		return
	}
	ctx.JSON(http.StatusOK, hub) // return hub details as json
}
