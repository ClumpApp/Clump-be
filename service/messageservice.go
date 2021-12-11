package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) GetGroupMessages(groupid float64) []model.MessageDTO {
	var messageDTOs []model.MessageDTO
	obj.db.Query(&model.Message{}, &model.Message{GroupID: uint(groupid)}, &messageDTOs)
	return messageDTOs
}

func (obj *Service) CreateMessage(groupid, userid float64, messageDTO model.MessageDTO) model.MessageDTO {
	var message model.Message
	utility.Convert(&messageDTO, &message)
	message.GroupID = uint(groupid)
	message.UserID = uint(userid)
	obj.db.Create(&model.Message{}, &message)
	var out model.MessageDTO
	utility.Convert(&message, &out)
	return out
}

func (obj *Service) UpdateMessage(id string, messageDTO model.MessageDTO) {
	var message model.Message
	utility.Convert(&messageDTO, &message)
	message.MessageEdited = true
	uuid := utility.ConvertUUID(id)
	obj.db.Update(&model.Message{}, &model.Message{UUID: uuid}, &message)
}

func (obj *Service) DeleteMessage(id string) {
	uuid := utility.ConvertUUID(id)
	obj.db.Delete(&model.Message{}, &model.Message{UUID: uuid})
}
