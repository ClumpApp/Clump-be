package service

import (
	"github.com/clumpapp/clump-be/database"
)

type Service struct {
	db       *database.Database
	delegate *MessageDelegate
}

func New(db *database.Database) *Service {
	return &Service{db: db}
}

func (obj *Service) SetDelegate(messageDelegate MessageDelegate) {
	obj.delegate = &messageDelegate
}
