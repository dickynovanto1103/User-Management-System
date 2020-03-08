package main

import (
	"context"
	"log"
	"net"

	"github.com/dickynovanto1103/User-Management-System/container"
	tcpServer "github.com/dickynovanto1103/User-Management-System/internal/service/tcpserver"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/profile"

	pb "github.com/dickynovanto1103/User-Management-System/proto"
)

type server struct {
	pb.UnimplementedUserDataServiceServer
	repository *container.Repository
}

func (s *server) SendRequest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	requestID := in.GetRequestID()
	mapper := in.GetMapper()
	response := tcpServer.HandleRequest(requestID, mapper, s.repository)
	newResponse := &pb.Response{
		ResponseID:           response.ResponseID,
		Mapper:               response.Data,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return newResponse, nil
}

func main() {

	defer profile.Start().Stop()
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Println("error found in listening: ", err)
	}

	repository := container.BuildTCPServerDep()

	defer repository.DBImpl.CloseDB()

	grpcServer := grpc.NewServer()
	pb.RegisterUserDataServiceServer(grpcServer, &server{
		repository: repository,
	})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
