package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID           uuid.UUID `gorm:"type:uuid"`
	Name           string
	ProfilePicture string
	UserName       string
	UserMail       string
	Password       string
	GroupID        uint
	UserGroups     []IEUserGroup
	UserInterests  []IEUserInterest
	Messages       []Message
}

type Interest struct {
	gorm.Model
	UUID          uuid.UUID `gorm:"type:uuid"`
	Title         string
	Picture       string
	UserInterests []IEUserInterest
}

type IEUserInterest struct {
	gorm.Model
	UserID     uint
	InterestID uint
}

type Group struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"type:uuid"`
	Users      []User
	UserGroups []IEUserGroup
	IsOpen     bool
	Message    []Message
	Board      string
}

type IEUserGroup struct {
	gorm.Model
	UserID    uint
	GroupID   uint
	EntryDate time.Time
}

type Message struct {
	gorm.Model
	UUID          uuid.UUID `gorm:"type:uuid"`
	UserID        uint
	GroupID       uint
	MessageType   string
	MessageText   string
	MessageEdited bool
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = uuid.New()
	return
}

func (obj *Group) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = uuid.New()
	return
}

func (obj *Interest) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = uuid.New()
	return
}

func (obj *Message) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = uuid.New()
	return
}
