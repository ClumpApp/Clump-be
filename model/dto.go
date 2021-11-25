package model

type UserDTO struct {
	ID             uint
	Name           string
	ProfilePicture string
	UserName       string
	UserMail       string
	GroupID        uint
}

type InterestDTO struct {
	ID      uint
	Title   string
	Picture string
}

type GroupDTO struct {
	ID     uint
	IsOpen bool
	Board  string
}

type MessageDTO struct {
	ID            uint
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
