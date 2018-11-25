package v1

import (
	"fmt"
	"github.com/eiji03aero/thp_server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var (
		users []*models.User
		err   error
	)

	if users, err = models.GetUsers(); err != nil {
		c.String(http.StatusInternalServerError, "error occured")
		return
	}

	for _, user := range users {
		fmt.Println("user: ", user)
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
