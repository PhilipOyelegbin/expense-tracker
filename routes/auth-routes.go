package routes

import (
	"expense-tracker/controller"

	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/auth/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/auth/login", controller.LoginUser).Methods("POST")
}
