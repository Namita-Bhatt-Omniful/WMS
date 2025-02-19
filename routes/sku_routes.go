package routes

import (
	"WMS/controllers"

	"github.com/omniful/go_commons/http"
)

func SKUroutes(r *http.Server) {
	route := r.Group("/sku")
	{
		route.POST("/verify", controllers.VerifySKU)

	}
}
