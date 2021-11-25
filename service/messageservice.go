package service

import (
	"github.com/clumpapp/clump-be/model"
	"gorm.io/gorm"
)

func (obj *Service) UpdateMessage(messageDTO model.MessageDTO) {
	obj.db.Update(&model.Message{}, &model.User{Model: gorm.Model{ID: messageDTO.ID}}, &messageDTO)
	obj.db.Update(&model.Message{}, &model.User{Model: gorm.Model{ID: messageDTO.ID}}, &model.Message{MessageEdited: true})
}

func (obj *Service) DeleteMessage(messageDTO model.MessageDTO) {
	obj.db.Delete(&model.Message{}, &model.Message{Model: gorm.Model{ID: messageDTO.ID}})
}
