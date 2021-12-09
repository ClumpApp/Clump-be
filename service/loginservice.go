package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) Login(loginDTO model.LoginDTO) bool {
	var login model.LoginDTO
	obj.db.Query(&model.User{}, &model.User{UserName: loginDTO.UserName}, &login)
	if login.UserName != "" && login.Password != "" {
		return utility.CompareHash(loginDTO.Password, login.Password)
	}
	return false
}
