package main

import (
	"context"
	"fmt"
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
	doGreet(client)
	//doGreetManytimes(client)

	// router := gin.Default()

	// router.GET("/test", func(ctx *gin.Context) {

	// 	client := pb.NewGreetServiceClient(conn)
	// 	fmt.Println(client)

	// 	ctx.JSON(200, DoLongGreet(client))

	// })
	// router.Run(":5000")

}

func DoLongGreet(ctx pb.GreetServiceClient) string {

	reqs := []*pb.GreetRequest{
		{FirstName: "Et Bil2u"},
		{FirstName: "asdf"},
		{FirstName: "Hello"},
		{FirstName: "World"},
	}

	stream, _ := ctx.LongGreet(context.Background())

	for _, req := range reqs {
		stream.Send(req)
		//time.Sleep(time.Second * 1)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("Error :", err.Error())
	}

	return resp.Result

}
