package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) Login(loginDTO model.LoginDTO) bool {
	var login model.LoginDTO
	obj.db.Query(&model.User{}, &model.User{UserName: loginDTO.UserName}, &login)
	result := utility.CompareHash(loginDTO.Password, login.Password)
	return result
}
