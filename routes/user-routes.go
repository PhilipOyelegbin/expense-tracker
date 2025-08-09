package routes

import (
	"expense-tracker/controller"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/me", controller.GetMyAccount).Methods("GET")
	router.HandleFunc("/users/me", controller.DeleteMyAccount).Methods("DELETE")
}