package main

import (
	"context"
	"fmt"
	pb "golang-grpc/greet/proto"
	"time"
)

func doLongGreet(ctx pb.GreetServiceClient) {

	reqs := []*pb.GreetRequest{
		{FirstName: "Et Bilu"},
		{FirstName: "asdf"},
		{FirstName: "Hello"},
		{FirstName: "World"},
	}

	stream, _ := ctx.LongGreet(context.Background())

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 1)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("Error :", err.Error())
	}

	fmt.Println("LongGreet -> ", resp.Result)

}
