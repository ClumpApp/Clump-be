package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) UpdateGroup(id string, groupDTO model.GroupDTO) {
	var group model.Group
	utility.Convert(&groupDTO, &group)
	uuid := utility.ConvertUUID(id)
	obj.db.Update(&model.Group{}, model.Group{UUID: uuid}, &group)
}

func (obj *Service) DeleteGroup(id string) {
	uuid := utility.ConvertUUID(id)
	obj.db.Delete(&model.Group{}, model.Group{UUID: uuid})
}

/*
func (obj *Service) UpdateOpen(groupDTO model.GroupDTO) {
	//obj.db.Update(&model.GroupDTO{}, &model.GroupDTO{ID: groupDTO.ID}, &model.GroupDTO{IsOpen: isOpen})
	obj.db.Update(&model.Group{}, &model.Group{Model: gorm.Model{ID: groupDTO.ID}}, &model.Group{IsOpen: !groupDTO.IsOpen})
}

func (obj *Service) UpdateBoard(groupDTO model.GroupDTO, board string) {
	obj.db.Update(&model.Group{}, &model.Group{Model: gorm.Model{ID: groupDTO.ID}}, &model.Group{Board: board})
}
*/
