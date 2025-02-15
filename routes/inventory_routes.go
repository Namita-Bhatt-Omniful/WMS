package routes

import (
	"WMS/controllers"

	"github.com/omniful/go_commons/http"
)

func InventoryRoutes(r *http.Server) {
	route := r.Group("/inventory")
	{
		route.GET("/view", controllers.GetInventory)
		route.POST("/update", controllers.UpdateInventory)
	}
}
