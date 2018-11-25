package scripts

import (
	"github.com/eiji03aero/thp_server/models"
)

func Migrate() {
	db := models.Setup()

	db.AutoMigrate(
		&models.User{},
		&models.Article{},
	)
}
