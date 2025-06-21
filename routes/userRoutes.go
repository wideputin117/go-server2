package routes

import (
	"net/http"

	"example.com/go-server/controllers"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.RegisterUser(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
