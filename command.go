package main

import (
	"flag"
	"fmt"
	_ "github.com/eiji03aero/thp_server/models"
	"github.com/eiji03aero/thp_server/scripts"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "db:reset":
		scripts.ResetDB()

	case "db:migrate":
		scripts.Migrate()

	case "db:seed":
		scripts.Seed()

	default:
		fmt.Println("Ain't got nothing from first argument", flag.Arg(0))
	}
}
