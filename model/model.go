package model

import (
	"time"

	"github.com/clumpapp/clump-be/utility"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID           [16]byte `gorm:"type:uuid"`
	UserName       string
	UserMail       string
	Password       string
	ProfilePicture string
	GroupID        uint
	UserGroups     []IEUserGroup
	UserInterests  []IEUserInterest
	Messages       []Message
}

type Group struct {
	gorm.Model
	UUID           [16]byte `gorm:"type:uuid"`
	Users          []User
	UserGroups     []IEUserGroup
	Messages       []Message
	GroupInterests []IEGroupInterest
}

type IEUserGroup struct {
	gorm.Model
	UserID    uint
	GroupID   uint
	EntryDate time.Time
}

type Interest struct {
	gorm.Model
	UUID             [16]byte `gorm:"type:uuid"`
	Title            string
	Picture          string
	UserInterests    []IEUserInterest
	GroupInterests   []IEGroupInterest
	SubInterests     []Interest `gorm:"foreignkey:InterestID"`
	InterestID       *uint
	SubInterestCount uint
}

type IEUserInterest struct {
	gorm.Model
	UserID     uint
	InterestID uint
}

type IEGroupInterest struct {
	gorm.Model
	GroupID    uint
	InterestID uint
}

type MessageType int

const (
	Undefined MessageType = iota
	Text
	Image
	Video
	Other
)

type Message struct {
	gorm.Model
	UUID          [16]byte `gorm:"type:uuid"`
	UserID        uint
	User          User
	GroupID       uint
	MessageType   MessageType
	MessageString string
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = utility.NewUUID()
	return
}

func (obj *Group) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = utility.NewUUID()
	return
}

func (obj *Interest) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = utility.NewUUID()
	return
}

func (obj *Message) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = utility.NewUUID()
	return
}

func (obj *Message) AfterFind(tx *gorm.DB) (err error) {
	if obj.MessageType == Image || obj.MessageType == Video {
		obj.MessageString = utility.GetStorage().GetURL() + obj.MessageString
	}
	return
}
