package model

type UserDTO struct {
	ID             uint `json:"omitempty"`
	Name           string
	ProfilePicture string
	UserName       string
	UserMail       string
	GroupID        uint
}

type InterestDTO struct {
	ID      uint `json:"omitempty"`
	Title   string
	Picture string
}

type GroupDTO struct {
	ID     uint `json:"omitempty"`
	IsOpen bool
	Board  string
}

type MessageDTO struct {
	ID            uint `json:"omitempty"`
	UserID        uint
	GroupID       uint
	MessageType   string
	MessageText   string
	MessageEdited bool
}

type LoginDTO struct {
	UserName string
	Password string
}
