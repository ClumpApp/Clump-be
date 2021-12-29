package service

import (
	"io"

	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) GetGroupMessages(groupid float64) []model.MessageOutDTO {
	var messages []model.Message
	obj.db.QueryWithPreload(&model.Message{}, &model.Message{GroupID: uint(groupid)}, &messages)
	var messageDTOs []model.MessageOutDTO
	for _, meesage := range messages {
		messageDTOs = append(messageDTOs, model.MessageOutDTO{
			UUID:          meesage.UUID.String(),
			UserName:      meesage.User.UserName,
			MessageType:   int(meesage.MessageType),
			MessageString: meesage.MessageString,
			CreatedAt:     meesage.CreatedAt,
		})
	}
	return messageDTOs
}

func (obj *Service) createMessage(groupid, userid float64, messageString string, messageType model.MessageType) {
	message := model.Message{
		GroupID:       uint(groupid),
		UserID:        uint(userid),
		MessageType:   messageType,
		MessageString: messageString,
	}
	obj.db.Create(&model.Message{}, &message)
}

func (obj *Service) createMedia(groupid, userid float64, name string, file io.ReadSeekCloser, messageType model.MessageType) {
	newName := obj.uploadMedia(name, file)
	obj.createMessage(groupid, userid, newName, messageType)
}

func (obj *Service) CreateMessage(groupid, userid float64, messageDTO model.MessageInDTO) {
	obj.createMessage(groupid, userid, messageDTO.Message, model.Text)
}

func (obj *Service) CreateImage(groupid, userid float64, name string, file io.ReadSeekCloser) {
	obj.createMedia(groupid, userid, name, file, model.Image)
}

func (obj *Service) CreateVideo(groupid, userid float64, name string, file io.ReadSeekCloser) {
	obj.createMedia(groupid, userid, name, file, model.Video)
}

func (obj *Service) CreateOther(groupid, userid float64, name string, file io.ReadSeekCloser) {
	obj.createMedia(groupid, userid, name, file, model.Other)
}

func (obj *Service) DeleteMessage(uuid string, userid float64) {
	uuID := utility.ConvertUUID(uuid)
	userID := uint(userid)
	var message model.Message
	found := obj.db.Read(&model.Message{}, &model.Message{UUID: uuID, UserID: userID}, &message)
	if found {
		if message.MessageType == model.Image || message.MessageType == model.Video || message.MessageType == model.Other {
			obj.deleteMedia(message.MessageString)
		}
		obj.db.Delete(&model.Message{}, message.ID)
	}
}
