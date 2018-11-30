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

func GetUser(id int) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser(user *User) error {
	err := db.Save(user).Error
	return err
}

func EditUser(id int, user *User) error {
	if err := db.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}
	return nil
}
