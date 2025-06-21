package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go-server/database"
	"example.com/go-server/routes"
)

func main() {
	// Initialize MongoDB connection
	database.ConnectDB()

	// Create a new router (multiplexer)
	mux := http.NewServeMux()

	// Register routes
	routes.RegisterProductRoutes(mux)
	routes.RegisterUserRoutes(mux)
	// Start HTTP server
	fmt.Println("🚀 Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
