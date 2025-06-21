package routes

import (
	"net/http"

	"example.com/go-server/controllers"
)

// RegisterProductRoutes registers product-related routes
func RegisterProductRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.CreateProduct(w, r) // Create a new product
		case http.MethodGet:
			controllers.GetProducts(w, r) // Get all products
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
