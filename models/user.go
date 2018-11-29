package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:name`
	Email    string    `json:email`
	Articles []Article `json:articles`
}

func GetUsers() ([]*User, error) {
	var users []*User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func SaveUser(user *User) error {
	err := db.Save(user).Error
	return err
}
