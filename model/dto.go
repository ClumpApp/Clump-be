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

type MessageInDTO struct {
	Message string `json:"message"`
}

type MessageOutDTO struct {
	UUID          string    `json:"uuid"`
	UserName      string    `json:"username"`
	MessageType   int       `json:"type"`
	MessageString string    `json:"messagestr"`
	CreatedAt     time.Time `json:"date"`
}

type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
