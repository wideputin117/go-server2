package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go-server/database"
	"example.com/go-server/routes"
	"github.com/gorilla/mux" // ✅ Use Gorilla Mux
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()

	// Create a new router (multiplexer)
	router := mux.NewRouter()

	// Register routes
	routes.RegisterProductRoutes(router)
	routes.RegisterUserRoutes(router)
	// Start HTTP server
	fmt.Println("🚀 Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
