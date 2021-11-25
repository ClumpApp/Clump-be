package service

import (
	"github.com/clumpapp/clump-be/model"
	"gorm.io/gorm"
)

func (obj *Service) UpdateUser(userDTO model.UserDTO) {
	obj.db.Update(&model.User{}, &model.User{Model: gorm.Model{ID: userDTO.ID}}, &userDTO)
}

func (obj *Service) DeleteUser(userDTO model.UserDTO) {
	obj.db.Delete(&model.User{}, &model.User{Model: gorm.Model{ID: userDTO.ID}})
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
