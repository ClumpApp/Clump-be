package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) GetMessages(groupDTO model.GroupDTO) []model.MessageDTO {
	var messagesDTO []model.MessageDTO
	obj.db.Query(&model.Message{}, &model.Message{GroupID: groupDTO.ID}, &messagesDTO)
	return messagesDTO
}

func (obj *Service) GetUsers(groupDTO model.GroupDTO) []model.UserDTO {
	var usersDTO []model.UserDTO
	//var group model.GroupDTO
	//utility.Convert(&groupDTO, &group)
	obj.db.Query(&model.User{}, &model.User{GroupID: groupDTO.ID}, &usersDTO)
	return usersDTO
}

func (obj *Service) UpdateGroup(groupDTO model.GroupDTO) {
	var group model.User
	utility.Convert(&groupDTO, &group)
	obj.db.Update(&model.Group{}, groupDTO.ID, &group)
}

func (obj *Service) DeleteGroup(groupDTO model.GroupDTO) {
	obj.db.Delete(&model.Group{}, groupDTO.ID)
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
