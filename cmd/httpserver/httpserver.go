package main

import (
	"log"
	"net/http"

	"github.com/dickynovanto1103/User-Management-System/container"
	"github.com/dickynovanto1103/User-Management-System/internal/service/httpserver"
	"google.golang.org/grpc"
)

func main() {
	httpserver.HandleRouting()

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error dialing grpc: ", err)
	}
	defer conn.Close()

	container.BuildHttpServerClient(conn)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
