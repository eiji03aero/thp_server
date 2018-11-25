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

func GetUsers() ([]*User, error) {
	var users []*User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
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
