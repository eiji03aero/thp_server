package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Setup() *gorm.DB {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=eijiosakabe dbname=thp sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		panic(err.Error())
	}
}
