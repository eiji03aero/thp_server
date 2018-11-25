package scripts

import (
	"github.com/eiji03aero/thp_server/models"
)

func Seed() {
	db := models.Setup()

	models.CreateUser(map[string]interface{}{
		"name":  "eijiosakabe",
		"email": "eiji03aero@gmail.com",
	})

	var user models.User
	db.First(&user)

	models.CreateArticle(map[string]interface{}{
		"title":   "Tips on dev",
		"body":    "Just do it man",
		"user_id": user.ID,
	})
	models.CreateArticle(map[string]interface{}{
		"title":   "second article",
		"body":    "kore ga body",
		"user_id": user.ID,
	})
	models.CreateArticle(map[string]interface{}{
		"title":   "second article",
		"body":    "kore ga body",
		"user_id": user.ID,
	})
}
