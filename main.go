package main

import (
	appinit "WMS/init"
	"WMS/routes"
	"fmt"

	"github.com/omniful/go_commons/http"
)

func main() {
	server := http.InitializeServer(":8080", 0, 0, 0)
	appinit.InitializeDB()
	routes.HubRoutes(server)
	routes.InventoryRoutes(server)
	err1 := server.StartServer("WMS")

	if err1 != nil {
		fmt.Println("Failed to start the server:", err1)
		return
	}
	fmt.Println("Server started successfully!")
}
