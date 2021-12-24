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

func (obj *Service) CreateMessage(groupid, userid float64, messageDTO model.MessageInDTO) {
	var message model.Message
	utility.Convert(&messageDTO, &message)
	message.GroupID = uint(groupid)
	message.UserID = uint(userid)
	message.MessageType = model.Text
	obj.db.Create(&model.Message{}, &message)
}

func (obj *Service) createMedia(groupid, userid float64, name string, file io.ReadSeekCloser, mType model.MessageType) {
	newName := obj.uploadMedia(name, file)
	message := model.Message{
		GroupID:       uint(groupid),
		UserID:        uint(userid),
		MessageType:   mType,
		MessageString: newName,
	}
	obj.db.Create(&model.Message{}, &message)
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

func (obj *Service) DeleteMessage(id string) {
	uuid := utility.ConvertUUID(id)
	var message model.Message
	obj.db.Read(&model.Message{}, &model.Message{UUID: uuid}, &message)
	if message.MessageType == model.Image || message.MessageType == model.Video {
		obj.deleteMedia(message.MessageString)
	}
	obj.db.Delete(&model.Message{}, &model.Message{UUID: uuid})
}
