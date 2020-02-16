package main

import (
	"context"
	"github.com/dickynovanto1103/User-Management-System/internal/service/tcpServer"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/dickynovanto1103/User-Management-System/internal/redisutil"
	"github.com/dickynovanto1103/User-Management-System/internal/service/config"

	"github.com/dickynovanto1103/User-Management-System/internal/dbutil"

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
	response := tcpServer.HandleRequest(requestID, mapper)
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
	configDB := config.LoadConfigDB("config/configDB.json")
	configRedis := config.LoadConfigRedis("config/configRedis.json")

	redisutil.CreateRedisClient(configRedis)

	defer profile.Start().Stop()
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Println("error found in listening: ", err)
	}

	dbutil.PrepareDB(configDB)
	defer dbutil.CloseDB()
	dbutil.PrepareStatements()

	grpcServer := grpc.NewServer()
	pb.RegisterUserDataServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
