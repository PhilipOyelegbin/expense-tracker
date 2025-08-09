package routes

import (
	"expense-tracker/controller"

	"github.com/gorilla/mux"
)

var RegisterExpenseRoutes = func(router *mux.Router) {
	router.HandleFunc("/expenses", controller.CreateExpense).Methods("POST")
	router.HandleFunc("/expenses", controller.GetExpense).Methods("GET")
	router.HandleFunc("/expenses/{id}", controller.GetExpenseById).Methods("GET")
	router.HandleFunc("/expenses/{id}", controller.UpdateExpense).Methods("PATCH")
	router.HandleFunc("/expenses/{id}", controller.DeleteExpenseById).Methods("DELETE")
}