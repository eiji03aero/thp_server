package main

import (
	"github.com/eiji03aero/thp_server/models"
	"github.com/eiji03aero/thp_server/routers"
)

func main() {
	models.Setup()

	r := routers.InitRouter()

	r.Run(":8080")
}
