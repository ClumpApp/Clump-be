package service

import (
	"github.com/clumpapp/clump-be/database"
)

type Service struct {
	db *database.Database
}

func New(db *database.Database) *Service {
	return &Service{db}
}
