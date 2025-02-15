package routes

import (
	"WMS/controllers"

	"github.com/omniful/go_commons/http"
)

func HubRoutes(r *http.Server) {
	route := r.Group("/hubs")
	{
		route.GET("/:id", controllers.GetHubByID)
	}
}
