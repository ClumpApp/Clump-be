package model

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	UserID     int
	Username   string
	ChatroomID int
	EnteryDate time.Time
}

type Chatroom struct {
	gorm.Model
	ChatroomID   int
	Name         string
	Participants []int
	IsOpen       bool
	Interests    []string
}

type UserAccount struct {
	gorm.Model
	UserID       int
	Username     string
	Password     int
	PasswordHash int
	UserMail     string
	PasswordSalt string
	AccountType  string
}

type Message struct {
	gorm.Model
	MessageID   int
	UserID      int
	ChatroomID  int
	MessageType string
	Message     string
	SentDate    time.Time
}
