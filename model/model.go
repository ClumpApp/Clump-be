package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	ProfilePicture string
	Interest       []Interest
	GroupID        uint
	UserGroups     []UserGroups
	Message        []Message
}

type Account struct {
	gorm.Model
	UserID       uint
	UserName     string
	UserMail     string
	PasswordHash string
	PasswordSalt string
}

type Interest struct {
	gorm.Model
	UserID  uint
	Title   string
	Picture string
}

type Group struct {
	gorm.Model
	Users      []User
	UserGroups []UserGroups
	IsOpen     bool
	Message    []Message
}

type UserGroups struct {
	gorm.Model
	UserID    uint
	GroupID   uint
	EntryDate time.Time
}

type Message struct {
	gorm.Model
	UserID      uint
	GroupID     uint
	MessageType string
	Message     string
	SentDate    time.Time
}
