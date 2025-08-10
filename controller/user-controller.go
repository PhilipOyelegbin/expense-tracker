package controller

import (
	"encoding/json"
	"expense-tracker/model"
	"expense-tracker/utils"
	"net/http"
)

// User struct to represent a user in the API
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// @Tags User
// @Summary Get my profile
// @Description Get my profile as a signed in user
// @Accept  json
// @Produce json
// @Success 200 {object} User "Successful operation"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/me [get]
func GetMyAccount(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

	user, _ := model.GetUserById(userId)
	if user.ID == 0 {
		http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Tags User
// @Summary Delete my profile
// @Description Delete my profile as a signed in user
// @Accept  json
// @Produce json
// @Success 204 {string} string "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/me [delete]
func DeleteMyAccount(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromJWTToken(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

	// Check if user exists
    user, _ := model.GetUserById(userId)
    if user.ID == 0 {
        http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
        return
    }

	deletedUser := model.DeleteUserById(userId)
	if deletedUser.ID == 0 {
		http.Error(w, `{"message": "Failed to delete user"}`, http.StatusInternalServerError)
        return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
