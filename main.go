package main

import (
	"clump/api"
	"clump/database"
	"clump/model"
	"clump/service"
	"clump/utility"
)

func main() {
	db := database.New()
	db.Connect()
	db.Migrate()

	user := model.User{UserName: "srr", Password: utility.GetHash("123456")}
	db.Create(&model.User{}, &user)

	s := service.New(db)
	a := api.New(s)

	a.Run()

}
