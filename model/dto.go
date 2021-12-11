package model

// IDs should not be exposed to client, UUID should be used

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

type GroupDTO struct {
	UUID   string `json:",omitempty"`
	IsOpen bool
	Board  string
}

type MessageDTO struct {
	UUID          string `json:",omitempty"`
	MessageType   string
	MessageText   string
	MessageEdited bool
}

type LoginDTO struct {
	UserName string
	Password string
}
