package service

import (
	"github.com/clumpapp/clump-be/model"
)

func (obj *Service) TextShare(messageDTO model.MessageDTO) {
	obj.db.Create(&model.Message{}, &messageDTO)
}
