package main

import (
	pb "golang-grpc/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listen server: %v", err)
	}

	log.Printf("Listenin on %s \n", addr)

	server := grpc.NewServer()

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Erorr serve: %v", err)
	}

}
