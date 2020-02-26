package main

import (
	"context"
	"log"
	"net"

	"github.com/dickynovanto1103/User-Management-System/container"
	tcpserver "github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/profile"

	pb "github.com/dickynovanto1103/User-Management-System/proto"
)

type server struct {
	pb.UnimplementedUserDataServiceServer
}

func (s *server) SendRequest(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	requestID := in.GetRequestID()
	mapper := in.GetMapper()
	response := tcpserver.HandleRequest(requestID, mapper)
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

	container.BuildTCPServerDep()

	defer container.DBImpl.CloseDB()

	grpcServer := grpc.NewServer()
	pb.RegisterUserDataServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
