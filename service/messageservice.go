package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) UpdateMessage(messageDTO model.MessageDTO) {
	var message model.User
	utility.Convert(&messageDTO, &message)
	obj.db.Update(&model.Message{}, messageDTO.ID, &messageDTO)
	obj.db.Update(&model.Message{}, messageDTO.ID, &model.Message{MessageEdited: true})
}

func (obj *Service) DeleteMessage(messageDTO model.MessageDTO) {
	obj.db.Delete(&model.Message{}, messageDTO.ID)
}
