package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) Login(loginDTO model.LoginDTO) (bool, model.UserDTO) {
	var user model.User
	found := obj.db.Read(&model.User{}, &model.User{UserName: loginDTO.UserName}, &user)
	if found && utility.CompareHash(loginDTO.Password, user.Password) {
		var out model.UserDTO
		utility.Convert(&user, &out)
		return true, out
	}
	return false, model.UserDTO{}
}

//this version doesnt have interests (will be updated)
func (obj *Service) SignUp(userDTO model.UserDTO) {
	var user model.User
	utility.Convert(&userDTO, &user)
	obj.db.Create(&model.User{}, &user)
}

func (obj *Service) GetGroupUsers(groupid string) []model.UserDTO {
	uid := utility.ConvertID(groupid)
	var usersDTO []model.UserDTO
	obj.db.Query(&model.User{}, &model.User{GroupID: uid}, &usersDTO)
	return usersDTO
}

func (obj *Service) UpdateUser(id string, userDTO model.UserDTO) {
	var user model.User
	utility.Convert(&userDTO, &user)
	obj.db.Update(&model.User{}, id, &user)
}

func (obj *Service) DeleteUser(id string) {
	obj.db.Delete(&model.User{}, id)
}

/* These are unnecessary as we will only be taking DTO struct from the request and can do them all once
func (obj *Service) UpdateUserName(userDTO model.UserDTO, name string) {
	//obj.db.Update(&model.UserDTO{}, &model.UserDTO{ID: userDTO.ID}, &model.UserDTO{UserName: name})
	//obj.db.Update(&model.User{}, &model.User{Model: gorm.Model{ID: userDTO.ID}}, &model.User{UserName: name})

}

func (obj *Service) UpdateProfilePicture(userDTO model.UserDTO, picture string) {
	//obj.db.Update(&model.UserDTO{}, &model.UserDTO{ID: userDTO.ID}, &model.UserDTO{ProfilePicture: picture})
	//obj.db.Update(&model.User{}, &model.User{Model: gorm.Model{ID: userDTO.ID}}, &model.User{ProfilePicture: picture})
}
*/
