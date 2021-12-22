package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) CreateMedia(groupid, userid float64, name string, data []byte) model.MessageDTO {
	var message model.Message
	message.GroupID = uint(groupid)
	message.UserID = uint(userid)
	// Determine file format, then assign a 16 char random name and upload
	obj.db.Create(&model.Message{}, &message)
	var out model.MessageDTO
	utility.Convert(&message, &out)
	out.MessageText = utility.GetStorage().GetURL() + out.MessageText
	return out
}

func (obj *Service) DeleteMedia(id string) {
	uuid := utility.ConvertUUID(id)
	var media model.Message
	obj.db.Read(&model.Message{}, &model.Message{UUID: uuid}, &media)
	utility.GetStorage().Delete(media.MessageText)
	obj.db.Delete(&model.Message{}, &model.Message{UUID: uuid})
}
