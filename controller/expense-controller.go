package controller

import (
	"encoding/json"
	"expense-tracker/model"
	"expense-tracker/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Expense struct to represent an expense in the API
type Expense struct {
	Title      string  `json:"title"`
	Description string  `json:"description"`
	Amount     float64 `json:"amount"`
	Date       string  `json:"date"`
	Category   string  `json:"category"`
}

// Categories struct to represent an expense category in the API
var Categories = [7]string{
	"Groceries",
	"Leisure",
	"Electronics",
	"Utilities",
	"Clothing",
	"Health",
	"Others",
}


// @Tags Expense
// @Summary Get all expenses
// @Description Retrieve a list of all expenses
// @Accept  json
// @Produce json
// @Success 200 {array} Expense "Successful operation"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses [get]
func GetExpense(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
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
        http.Error(w, `{"message": "No expenses found"}`, http.StatusNotFound)
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
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/{id} [get]
func GetExpenseById(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
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
// @Summary Filter expenses by past week
// @Description Retrieve a list of all expenses for the past week
// @Accept  json
// @Produce json
// @Success 200 {array} Expense "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/week [get]
func FilterExpenseByWeek(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
        return
    }

	// get the past week expenses based on the current day
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	var pastWeekExpenses []model.ExpenseData
	for _, expense := range model.GetExpense() {
		expenseDate, err := time.Parse("02/01/2006", expense.Date)
		if err != nil {
			log.Printf("Failed to parse date for expense ID %d: %v", expense.ID, err)
			continue
		}

		if expense.UserId == userId && expenseDate.After(sevenDaysAgo) {
			pastWeekExpenses = append(pastWeekExpenses, expense)
		}
	}

	if len(pastWeekExpenses) == 0 {
		http.Error(w, `{"message": "No expenses found for the past week"}`, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(pastWeekExpenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// @Tags Expense
// @Summary Filter expenses by past month
// @Description Retrieve a list of all expenses for the past month
// @Accept  json
// @Produce json
// @Success 200 {array} Expense "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/month [get]
func FilterExpenseByMonth(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
        return
    }

	// get the past month expenses based on the current day
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	var pastMonthExpenses []model.ExpenseData
	for _, expense := range model.GetExpense() {
		expenseDate, err := time.Parse("02/01/2006", expense.Date)
		if err != nil {
			log.Printf("Failed to parse date for expense ID %d: %v", expense.ID, err)
			continue
		}

		if expense.UserId == userId && expenseDate.After(oneMonthAgo) {
			pastMonthExpenses = append(pastMonthExpenses, expense)
		}
	}

	if len(pastMonthExpenses) == 0 {
		http.Error(w, `{"message": "No expenses found for the past month"}`, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(pastMonthExpenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// @Tags Expense
// @Summary Filter expenses by past three month
// @Description Retrieve a list of all expenses for the past three month
// @Accept  json
// @Produce json
// @Success 200 {array} Expense "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/past-three-month [get]
func FilterExpenseByPastThreeMonth(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
        return
    }

	// get the past three month expenses based on the current day
	threeMonthAgo := time.Now().AddDate(0, -3, 0)
	var pastThreeMonthExpenses []model.ExpenseData
	for _, expense := range model.GetExpense() {
		expenseDate, err := time.Parse("02/01/2006", expense.Date)
		if err != nil {
			log.Printf("Failed to parse date for expense ID %d: %v", expense.ID, err)
			continue
		}

		if expense.UserId == userId && expenseDate.After(threeMonthAgo) {
			pastThreeMonthExpenses = append(pastThreeMonthExpenses, expense)
		}
	}

	if len(pastThreeMonthExpenses) == 0 {
		http.Error(w, `{"message": "No expenses found for the past three month"}`, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(pastThreeMonthExpenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// @Tags Expense
// @Summary Filter expenses by custom date
// @Description Retrieve a list of all expenses by custom date
// @Accept  json
// @Produce json
// @Param start_date query string true "Start date in YYYY-MM-DD format"
// @Param end_date query string true "End date in YYYY-MM-DD format"
// @Success 200 {array} Expense "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /expenses/dates [get]
func FilterExpenseByCustomDate(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, _ := model.GetUserById(userId)
	if user.ID == 0 {
		http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// get the start and end date from query parameters.
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	if startDateStr == "" || endDateStr == "" {
		http.Error(w, `{"message": "Both start_date and end_date query parameters are required."}`, http.StatusBadRequest)
		return
	}

	// parse the date strings into time.Time objects.
	const dateFormat = "02/01/2006"
	startDate, err := time.Parse(dateFormat, startDateStr)
	if err != nil {
		http.Error(w, `{"message": "Invalid start_date format. Please use YYYY-MM-DD."}`, http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse(dateFormat, endDateStr)
	if err != nil {
		http.Error(w, `{"message": "Invalid end_date format. Please use YYYY-MM-DD."}`, http.StatusBadRequest)
		return
	}

	var filteredExpenses []model.ExpenseData
	allExpenses := model.GetExpense()

	for _, expense := range allExpenses {
		// skip invalid expenses or those that don't belong to the user.
		if expense.UserId != userId {
			continue
		}

		expenseDate, err := time.Parse(dateFormat, expense.Date)
		if err != nil {
			log.Printf("Failed to parse date for expense ID %d: %v", expense.ID, err)
			continue
		}

		// check if the expense date is within the custom date range (inclusive).
		isAfterOrEqualStart := !expenseDate.Before(startDate)
		isBeforeOrEqualEnd := !expenseDate.After(endDate)
		if isAfterOrEqualStart && isBeforeOrEqualEnd {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}

	if len(filteredExpenses) == 0 {
		http.Error(w, `{"message": "No expenses found for the specified date range."}`, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(filteredExpenses)
	if err != nil {
		http.Error(w, `{"message": "Failed to marshal expenses."}`, http.StatusInternalServerError)
		return
	}
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
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
        return
    }

	// parse the request body to create a new expense
	newExpense := &model.ExpenseData{}
	utils.ParseBody(r, newExpense)

	if newExpense.Title == "" || newExpense.Description == "" || newExpense.Amount <= 0 || newExpense.Date == "" {
		http.Error(w, `{"message":"All fields are required."}`, http.StatusBadRequest)
		return
	}

	if newExpense.Category != Categories[0] && newExpense.Category != Categories[1] && newExpense.Category != Categories[2] && newExpense.Category != Categories[3] && newExpense.Category != Categories[4] && newExpense.Category != Categories[5] && newExpense.Category != Categories[6] {
		http.Error(w, `{"message":"Invalid category provided"}`, http.StatusBadRequest)
		return
	}

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
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
        return
    }

	// parse the data from the body and convert the id parameter from the request
	updateExpense := &model.ExpenseData{}
	utils.ParseBody(r, updateExpense)
	vars := mux.Vars(r)
	expenseId := vars["id"]
	ID, err := strconv.ParseInt(expenseId, 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "Unauthorized"}`, http.StatusUnauthorized)
        return
    }

	// get the ID from the paramter and convert to integer
	vars := mux.Vars(r)
	expenseId := vars["id"]
	ID, err := strconv.ParseInt(expenseId, 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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