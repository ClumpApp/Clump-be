package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) GetGroupMessages(id string) []model.MessageDTO {
	uid := utility.ConvertID(id)
	var messagesDTO []model.MessageDTO
	obj.db.Query(&model.Message{}, &model.Message{GroupID: uid}, &messagesDTO)
	return messagesDTO
}

func (obj *Service) CreateMessage(messageDTO model.MessageDTO) {
	var message model.Message
	utility.Convert(&messageDTO, &message)
	obj.db.Create(&model.Message{}, &message)
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
