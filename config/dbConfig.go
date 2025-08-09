package config

import (
	"expense-tracker/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var DB_URL string

func Connect() {
	env := utils.LoadEnv()
	DB_URL = env.DBURL
	con, err := gorm.Open("mysql", DB_URL)
	if err != nil {
		panic(err)
	}

	db = con
}

func GetDB () *gorm.DB {
	return db
}

