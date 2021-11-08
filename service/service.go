package service

import (
	"clump/database"
)

type Service struct {
	db *database.Database
}

func New(db *database.Database) *Service {
	return &Service{db}
}
