package service

import (
	"github.com/clumpapp/clump-be/model"
)

type MessageDelegate interface {
	SendMessage(groupID uint, message model.MessageOutDTO)
}
