package container

import (
	pb "github.com/dickynovanto1103/User-Management-System/proto"
	"google.golang.org/grpc"
)

var Client pb.UserDataServiceClient

func BuildHttpServerClient(conn *grpc.ClientConn) {
	Client = pb.NewUserDataServiceClient(conn)
}
