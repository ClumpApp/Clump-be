package service

import (
	"clump/model"
)

func (obj *Service) TextShare(messageDTO model.MessageDTO) {
	obj.db.Create(&model.Message{}, &messageDTO)
}
