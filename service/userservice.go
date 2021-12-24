package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) Login(loginDTO model.LoginDTO) (uint, uint, bool) {
	var user model.User
	found := obj.db.Read(&model.User{}, &model.User{UserName: loginDTO.UserName}, &user)
	if found && utility.CompareHash(loginDTO.Password, user.Password) {
		return user.ID, user.GroupID, true
	}
	return 0, 0, false
}

//this version doesnt have interests (will be updated)
func (obj *Service) SignUp(userDTO model.UserDTO) {
	var user model.User
	utility.Convert(&userDTO, &user)
	obj.db.Create(&model.User{}, &user)
}

func (obj *Service) GetGroupUsers(groupid float64) []model.UserDTO {
	var userDTOs []model.UserDTO
	obj.db.Query(&model.User{}, &model.User{GroupID: uint(groupid)}, &userDTOs)
	return userDTOs
}

func (obj *Service) GetUser(userid float64) model.UserDTO {
	var userDTO model.UserDTO
	obj.db.Query(&model.User{}, uint(userid), &userDTO)
	return userDTO
}

func (obj *Service) UpdateUser(id string, userDTO model.UserDTO) {
	var user model.User
	utility.Convert(&userDTO, &user)
	uuid := utility.ConvertUUID(id)
	obj.db.Update(&model.User{}, &model.User{UUID: uuid}, &user)
}

func (obj *Service) DeleteUser(id string) {
	uuid := utility.ConvertUUID(id)
	obj.db.Delete(&model.User{}, &model.User{UUID: uuid})
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
