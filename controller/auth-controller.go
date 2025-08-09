package controller

import (
	"encoding/json"
	"expense-tracker/model"
	"expense-tracker/utils"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Tags Auth
// @Summary Register a user
// @Description Register a new user
// @Accept  json
// @Produce json
// @Param user body User true "User data"
// @Success 201 {object} User "Successful operation"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// parse the body contents and hash the password
	newUser := &model.UserData{}
	utils.ParseBody(r, newUser)
	hash, err := argon2id.CreateHash(newUser.Password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	newUser.Password = hash

	// create the new user and marshall the contents as a response
	user := newUser.CreateUser()
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// @Tags Auth
// @Summary Login as a user
// @Description Login with user credentials
// @Accept  json
// @Produce json
// @Param user body Login true "User data"
// @Success 200 {string} string "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser model.UserData
	utils.ParseBody(r, &loginUser)

	if loginUser.Email == "" || loginUser.Password == "" {
		http.Error(w, `{"message": "Email and password are required"}`, http.StatusBadRequest)
		return
	}

	user := model.GetUsers()

	for _, u := range user {
		if u.Email == loginUser.Email {
			match, err := argon2id.ComparePasswordAndHash(loginUser.Password, u.Password)
			if err != nil || !match {
				http.Error(w, `{"message": "Invalid email or password"}`, http.StatusUnauthorized)
				return
			}
			token, err := utils.SignJWTToken(int64(u.ID), u.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Login successful", "token": "` + token + `"}`))
			return
		}
	}
	http.Error(w, `{"message": "Account does not exist"}`, http.StatusNotFound)
}
