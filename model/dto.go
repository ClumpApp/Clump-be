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
}

type MessageDTO struct {
	ID          uint
	UserID      uint
	GroupID     uint
	MessageType string
	MessageText string
}

type LoginDTO struct {
	UserName string
	Password string
}
