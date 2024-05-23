package main

import (
	pb "golang-grpc/blog/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listen server: %v", err)
	}

	log.Printf("Listenin on %s \n", addr)

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Erorr serve: %v", err)
	}
}
