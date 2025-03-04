package routes

import (
	"net/http"

	"example.com/go-server/controllers"
)

// RegisterProductRoutes registers product-related routes
func RegisterProductRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateProduct(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
}
