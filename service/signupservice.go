package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

//this version doesnt have interests (will be updated)
func (obj *Service) SignUp(userDTO model.UserDTO) {
	var user model.User
	utility.Convert(&userDTO, &user)
	obj.db.Create(&model.User{}, &user)
}
