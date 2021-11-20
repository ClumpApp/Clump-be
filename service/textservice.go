package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) TextShare(messageDTO model.MessageDTO) {
	var message model.Message
	utility.Convert(&messageDTO, &message)
	obj.db.Create(&model.Message{}, &message)
}
