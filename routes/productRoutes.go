package routes

import (
	"example.com/go-server/controllers"
	"github.com/gorilla/mux" // ✅ Use Gorilla Mux
)

// RegisterProductRoutes registers product-related routes
func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
}
