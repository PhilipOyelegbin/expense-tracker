package controller

import (
	"encoding/json"
	"expense-tracker/model"
	"expense-tracker/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewExpense model.ExpenseData

// Expense struct to represent an expense in the API
type Expense struct {
	Title      string  `json:"title"`
	Description string  `json:"description"`
	Amount     float64 `json:"amount"`
	Date       string  `json:"date"`
	Category   string  `json:"category"`
	UserID     int64   `json:"userId"`
}


// @Tags Expense
// @Summary Get all expenses
// @Description Retrieve a list of all expenses
// @Accept  json
// @Produce json
// @Success 200 {array} Expense "Successful operation"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses [get]
func GetExpense(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	allExpenses := model.GetExpense()
    var userExpenses []model.ExpenseData
    
    for _, expense := range allExpenses {
        if expense.UserId == userId {
            userExpenses = append(userExpenses, expense)
        }
    }

    if len(userExpenses) == 0 {
        http.Error(w, "No expenses found", http.StatusNotFound)
        return
    }

	res, err := json.Marshal(userExpenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Tags Expense
// @Summary Get all expenses
// @Description Retrieve a list of all expenses
// @Accept  json
// @Produce json
// @Param id path string true "Expense ID"
// @Success 200 {array} Expense "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/{id} [get]
func GetExpenseById(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// get the id parameter from the request and convert to integer
	vars := mux.Vars(r)
	expenseId := vars["id"]
	ID, err := strconv.ParseInt(expenseId, 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get expense first, then check ownership
	expense, _ := model.GetExpenseById(ID)
    if expense.ID == 0 {
        http.Error(w, `{"message": "Expense not found"}`, http.StatusNotFound)
        return
    }

    // Check if the expense belongs to the current user
    if expense.UserId != userId {
        http.Error(w, `{"message": "Unauthorized access to expense"}`, http.StatusForbidden)
        return
    }

    res, _ := json.Marshal(expense)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// @Tags Expense
// @Summary Create a expense
// @Description Create a new expense
// @Accept  json
// @Produce json
// @Param Expense body Expense true "Expense data"
// @Success 201 {object} Expense "Successful operation"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses [post]
func CreateExpense(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// parse the request body to create a new expense
	newExpense := &model.ExpenseData{}
	utils.ParseBody(r, newExpense)

	// append the userId to the new expense data
	newExpense.UserId = userId
	expense := newExpense.CreateExpense()
	res, err := json.Marshal(expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// @Tags Expense
// @Summary Update an expense
// @Description Update an expense
// @Accept  json
// @Produce json
// @Param id path string true "Expense ID"
// @Param Expense body Expense true "Expense data"
// @Success 204 {object} Expense "Successful operation"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Expense not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/{id} [patch]
func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	// add authorization check
    userId, err := utils.GetUserIdFromJWTToken(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

	// parse the data from the body and convert the id parameter from the request
	updateExpense := &model.ExpenseData{}
	utils.ParseBody(r, updateExpense)
	vars := mux.Vars(r)
	expenseId := vars["id"]
	ID, err := strconv.ParseInt(expenseId, 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// take the details of the expense and verify it against the update details
	expense, db := model.GetExpenseById(ID)
	if expense.ID == 0 {
		http.Error(w, `{"message": "Expense not found"}`, http.StatusNotFound)
		return
	}

	// Check ownership
    if expense.UserId != userId {
        http.Error(w, `{"message": "Unauthorized access to expense"}`, http.StatusForbidden)
        return
    }

	// update the existing expense, not the updateExpense struct
    if updateExpense.Title != "" {
        expense.Title = updateExpense.Title
    }
    if updateExpense.Description != "" {
        expense.Description = updateExpense.Description
    }
    if updateExpense.Date != "" {
        expense.Date = updateExpense.Date
    }
    if updateExpense.Category != "" {
        expense.Category = updateExpense.Category
    }
    if updateExpense.Amount != 0 {
        expense.Amount = updateExpense.Amount
    }

	// save the updated details to the database and marshal the details for a response
	db.Save(&expense)
	res, err := json.Marshal(expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

// @Tags Expense
// @Summary Delete an expense
// @Description This endpoint deletes an expense by its ID.
// @Accept  json
// @Produce json
// @Param id path string true "Expense ID"
// @Success 204 {string} string "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Expense not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/{id} [delete]
func DeleteExpenseById(w http.ResponseWriter, r *http.Request) {
	// add authorization check
    userId, err := utils.GetUserIdFromJWTToken(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

	// get the ID from the paramter and convert to integer
	vars := mux.Vars(r)
	expenseId := vars["id"]
	ID, err := strconv.ParseInt(expenseId, 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if expense exists and user owns it before deleting
    expense, _ := model.GetExpenseById(ID)
    if expense.ID == 0 {
        http.Error(w, `{"message": "Expense not found"}`, http.StatusNotFound)
        return
    }

    if expense.UserId != userId {
        http.Error(w, `{"message": "Unauthorized access to expense"}`, http.StatusForbidden)
        return
    }

	deletedExpense := model.DeleteExpenseById(ID)
    if deletedExpense.ID == 0 {
        http.Error(w, `{"message": "Failed to delete expense"}`, http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}