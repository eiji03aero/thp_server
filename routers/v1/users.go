package v1

import (
	"fmt"
	"github.com/eiji03aero/thp_server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersHandler(r *gin.RouterGroup) {
	users_routes := r.Group("/users")
	{
		users_routes.GET("/", func(c *gin.Context) {
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
		})

		users_routes.POST("/", func(c *gin.Context) {
			var user models.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			models.CreateUser(&user)
			c.JSON(http.StatusOK, gin.H{"user": user})
		})
	}
}
