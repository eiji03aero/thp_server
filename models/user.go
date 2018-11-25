package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Articles []Article
}

func CreateUser(data map[string]interface{}) error {
	user := User{
		Name:  data["name"].(string),
		Email: data["email"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
