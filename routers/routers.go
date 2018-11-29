package routers

import (
	"github.com/eiji03aero/thp_server/routers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
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

	v1_routes := r.Group("/v1")
	{
		v1.UsersHandler(v1_routes)
	}

	return r
}
