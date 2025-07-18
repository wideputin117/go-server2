package routes

import (
	"github.com/gorilla/mux" // 

	"example.com/go-server/controllers"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user/signup", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/user/profile/{id}", controllers.GetUserDetails).Methods("GET")
}
