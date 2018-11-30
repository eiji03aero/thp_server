package v1

import (
	"fmt"
	"github.com/eiji03aero/thp_server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UsersHandler(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	{
		r.GET("/", func(c *gin.Context) {
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

		r.GET("/:id", func(c *gin.Context) {
			var (
				id   int
				user *models.User
				err  error
			)

			if id, err = strconv.Atoi(c.Param("id")); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if user, err = models.GetUser(id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user})
		})

		r.POST("/", func(c *gin.Context) {
			var user models.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			models.SaveUser(&user)
			c.JSON(http.StatusOK, gin.H{"user": user})
		})

		r.PATCH("/:id", func(c *gin.Context) {
			var (
				id   int
				user models.User
				err  error
			)

			if id, err = strconv.Atoi(c.Param("id")); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := models.EditUser(id, &user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user})
		})

		r.DELETE("/:id", func(c *gin.Context) {
			var (
				id  int
				err error
			)

			if id, err = strconv.Atoi(c.Param("id")); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := models.DeleteUser(id); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"deleted": id})
		})
	}
}
