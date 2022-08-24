package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dbInstance, err := gorm.Open("mysql", os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}

	db = dbInstance
}

func GetDB() *gorm.DB {
	return db
}
