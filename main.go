package main

import (
	"github.com/eiji03aero/thp_server/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func main() {
	models.Setup()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main site",
			"info":  "domo",
		})
	})

	r.GET("/search", func(c *gin.Context) {
		var keyword string

		keyword = c.Query("keyword")

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main site",
			"info":  keyword,
		})
	})

	r.Run(":8080")
}
