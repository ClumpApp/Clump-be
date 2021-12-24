package model

import (
	"time"
)

type UserDTO struct {
	UUID           string `json:",omitempty"`
	Name           string
	ProfilePicture string
	UserName       string
	UserMail       string
}

type InterestDTO struct {
	UUID    string `json:",omitempty"`
	Title   string
	Picture string
}

type MessageDTO struct {
	UUID          string    `json:"uuid,omitempty"`
	UserName      string    `json:"username"`
	MessageType   int       `json:"type"`
	MessageString string    `json:"messagestr"`
	MessageDate   time.Time `json:"date"`
}

type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
