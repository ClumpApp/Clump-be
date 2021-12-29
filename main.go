package main

import (
	"github.com/clumpapp/clump-be/api"
	"github.com/clumpapp/clump-be/database"
	"github.com/clumpapp/clump-be/service"
)

func main() {
	db := database.New()
	db.Connect()

	s := service.New(db)
	a := api.New(s)
	s.SetDelegate(a)

	a.Run()
}
