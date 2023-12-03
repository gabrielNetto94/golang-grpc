package main

import (
	"log"

	pb "golang-grpc/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error ", err.Error())
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	//doGreet(client)
	//doGreetManytimes(client)
	doLongGreet(client)
}
