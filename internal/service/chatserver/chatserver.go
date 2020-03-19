package chatserver

type ChatServer interface {
	SendMessage(receiverId int, message string)
}
