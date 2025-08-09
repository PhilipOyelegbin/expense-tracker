package model

import (
	"expense-tracker/config"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type UserData struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ExpenseData struct {
	gorm.Model
	Title       string  `gorm:"json:title"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date       string  `json:"date"`
	Category   string  `json:"category"`
	UserId    int64   `json:"userId"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.DB().SetConnMaxLifetime(10 * 60 * 1000)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&UserData{}, &ExpenseData{})
	log.Println("Connected to database successfully")

}

func (u *UserData) CreateUser() *UserData {
	db.NewRecord(u)
	db.Create(&u)
	return u
}


func GetUsers() []UserData {
	var users []UserData
	db.Find(&users)
	return users
}

func GetUserById(id int64) (*UserData, *gorm.DB) {
	var user UserData
	db = db.Where("ID=?", id).Find(&user)
	return &user, db
}

func DeleteUserById(id int64) UserData {
	var user UserData
	db.Where("ID=?", id).Delete(user)
	return user
}

func (e *ExpenseData) CreateExpense() *ExpenseData {
	db.NewRecord(e)
	db.Create(&e)
	return e
}

func GetExpense() []ExpenseData {
	var expense []ExpenseData
	db.Find(&expense)
	return expense
}

func GetExpenseById(id int64) (ExpenseData, *gorm.DB) {
	var expense ExpenseData
	db = db.Where("ID=?", id).Find(&expense)
	return expense, db
}

func DeleteExpenseById(id int64) ExpenseData {
	var expense ExpenseData
	db.Where("ID=?", id).Delete(expense)
	return expense
}
