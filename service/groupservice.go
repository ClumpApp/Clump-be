package service

import (
	"github.com/clumpapp/clump-be/model"
	"gorm.io/gorm"
)

func (obj *Service) GetMessages(groupDTO model.GroupDTO) model.MessageDTO {
	var messages model.MessageDTO
	obj.db.Query(&model.MessageDTO{}, &model.MessageDTO{GroupID: groupDTO.ID}, &messages)
	return messages
}

func (obj *Service) GetUsers(groupDTO model.GroupDTO) model.UserDTO {
	var users model.UserDTO
	obj.db.Query(&model.UserDTO{}, &model.UserDTO{GroupID: groupDTO.ID}, &users)
	return users
}

func (obj *Service) UpdateGroup(groupDTO model.GroupDTO) {
	obj.db.Update(&model.Group{}, &model.Group{Model: gorm.Model{ID: groupDTO.ID}}, &groupDTO)
}

func (obj *Service) DeleteGroup(groupDTO model.GroupDTO) {
	obj.db.Delete(&model.Group{}, &model.Group{Model: gorm.Model{ID: groupDTO.ID}})
}

/*
func (obj *Service) UpdateOpen(groupDTO model.GroupDTO) {
	//obj.db.Update(&model.GroupDTO{}, &model.GroupDTO{ID: groupDTO.ID}, &model.GroupDTO{IsOpen: isOpen})
	obj.db.Update(&model.Group{}, &model.Group{Model: gorm.Model{ID: groupDTO.ID}}, &model.Group{IsOpen: !groupDTO.IsOpen})
}

func (obj *Service) UpdateBoard(groupDTO model.GroupDTO, board string) {
	obj.db.Update(&model.Group{}, &model.Group{Model: gorm.Model{ID: groupDTO.ID}}, &model.Group{Board: board})
}
*/
