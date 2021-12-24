package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) GetGroupMessages(groupid float64) []model.MessageDTO {
	var messages []model.Message
	obj.db.QueryWithPreload(&model.Message{}, &model.Message{GroupID: uint(groupid)}, &messages)
	var messageDTOs []model.MessageDTO
	for _, meesage := range messages {
		messageDTOs = append(messageDTOs, model.MessageDTO{
			UUID:          meesage.UUID.String(),
			UserName:      meesage.User.UserName,
			MessageType:   int(meesage.MessageType),
			MessageString: meesage.MessageString,
			MessageDate:   meesage.MessageDate,
		})
	}
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

func (obj *Service) DeleteMessage(id string) {
	uuid := utility.ConvertUUID(id)
	obj.db.Delete(&model.Message{}, &model.Message{UUID: uuid})
}
