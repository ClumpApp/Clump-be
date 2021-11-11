package service

import (
	"clump/model"
	"clump/utility"
)

func (obj *Service) Login(loginDTO model.LoginDTO) bool {
	var login model.LoginDTO
	obj.db.Query(&model.User{}, &model.User{UserName: loginDTO.UserName}, &login)
	result := utility.CompareHash(loginDTO.Password, login.Password)
	return result
}
