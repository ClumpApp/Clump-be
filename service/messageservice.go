package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) GetGroupMessages(groupid string) []model.MessageDTO {
	uid := utility.ConvertID(groupid)
	var messagesDTO []model.MessageDTO
	obj.db.Query(&model.Message{}, &model.Message{GroupID: uid}, &messagesDTO)
	return messagesDTO
}

func (obj *Service) CreateMessage(group, user string, messageDTO model.MessageDTO) model.MessageDTO {
	var message model.Message
	utility.Convert(&messageDTO, &message)
	message.GroupID = utility.ConvertID(group)
	message.UserID = utility.ConvertID(user)
	obj.db.Create(&model.Message{}, &message)
	var out model.MessageDTO
	utility.Convert(&message, &out)
	return out
}

func (obj *Service) UpdateMessage(id string, messageDTO model.MessageDTO) {
	var message model.Message
	utility.Convert(&messageDTO, &message)
	message.MessageEdited = true
	obj.db.Update(&model.Message{}, id, &messageDTO)
}

func (obj *Service) DeleteMessage(id string) {
	obj.db.Delete(&model.Message{}, id)
}
