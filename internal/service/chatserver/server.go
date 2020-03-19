package chatserver

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SendMessage(receiverId int, message string) {
	//create struct message
	//put message inside kafka
}
