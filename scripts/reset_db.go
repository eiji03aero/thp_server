package scripts

import (
	"github.com/eiji03aero/thp_server/models"
)

func ResetDB() {
	db := models.Setup()

	db.DropTableIfExists(models.User{}, models.Article{})
	db.CreateTable(models.User{}, models.Article{})
}
