package main

import (
	"golang-grpc/configs"
	pb "golang-grpc/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}

	path := "/home/gabriel/me/golang-grpc/"

	if configs.GetEnv("SSL") == "TRUE" {
		certFile := path + "ssl/server.crt"
		keyFile := path + "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("fail load cert: %v", err.Error())
		}
		opts = append(opts, grpc.Creds(creds))
	}

	log.Printf("Listenin on %s \n", addr)

	server := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(server, &Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Erorr serve: %v", err)
	}

}
