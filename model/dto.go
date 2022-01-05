package model

import (
	"time"
)

type UserDTO struct {
	UserName       string
	UserMail       string
	ProfilePicture string
}

type InterestDTO struct {
	UUID    string `json:"uuid,omitempty"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
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

type SignUpDTO struct {
	UserName string `json:"username"`
	UserMail string `json:"email"`
	Password string `json:"password"`
}
