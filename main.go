package main

import (
	"clump/api"
	"clump/database"
	"clump/service"
)

func main() {
	db := database.New()
	db.Connect()

	s := service.New(db)
	a := api.New(s)

	a.Run()
}
