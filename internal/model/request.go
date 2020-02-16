package model

type Request struct {
	RequestID int
	Data      map[string]string
}

const (
	RequestLogin          = 1
	RequestUserInfo       = 2
	RequestUpdateNickname = 3
	RequestUpdateProfile  = 4
)