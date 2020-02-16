package main

import (
	"github.com/dickynovanto1103/User-Management-System/container"
	"github.com/dickynovanto1103/User-Management-System/internal/service/httpServer"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	httpServer.HandleRouting()

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error dialing grpc: ", err)
	}
	defer conn.Close()

	container.BuildHttpServerClient(conn)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
