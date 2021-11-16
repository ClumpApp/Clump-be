package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	ProfilePicture string
	UserName       string
	UserMail       string
	Password       string
	UserInterests  []UserInterests
	GroupID        uint
	UserGroups     []UserGroups
	Messages       []Message
}

type Interest struct {
	gorm.Model
	Title         string
	Picture       string
	UserInterests []UserInterests
}

type UserInterests struct {
	gorm.Model
	UserID     uint
	InterestID uint
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
	MessageText string
}
