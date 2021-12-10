package main

import (
	"github.com/clumpapp/clump-be/api"
	"github.com/clumpapp/clump-be/database"
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/service"
	"github.com/clumpapp/clump-be/utility"
)

func main() {
	db := database.New()
	db.Connect()
	db.Migrate()

	group := model.Group{}
	db.Create(&model.Group{}, &group)

	user := model.User{UserName: "srr", Password: utility.GetHash("123456"), GroupID: group.ID}
	db.Create(&model.User{}, &user)

	s := service.New(db)
	a := api.New(s)

	a.Run()

}
