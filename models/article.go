package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title  string
	Body   string
	UserId uint
}

func CreateArticle(data map[string]interface{}) error {
	article := Article{
		Title:  data["title"].(string),
		Body:   data["body"].(string),
		UserId: data["user_id"].(uint),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}
