package chatserver

import "fmt"

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

type Message struct {
	msg      string
	receiver int
}

func NewMessage(msg string, receiver int) *Message {
	return &Message{
		msg:      msg,
		receiver: receiver,
	}
}

func (s *Server) SendMessage(receiverId int, message string) {
	//create struct message
	msg := NewMessage(message, receiverId)
	//put message inside kafka
	fmt.Println("msg:", msg)
}
