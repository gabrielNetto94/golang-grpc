package main

import (
	"context"
	"fmt"
	pb "golang-grpc/greet/proto"
	"io"
	"log"
)

func doGreetManytimes(ctx pb.GreetServiceClient) {

	fmt.Println("doGreetManytimes")
	req := &pb.GreetRequest{
		FirstName: "Gabriel",
	}

	resp, err := ctx.GreetManyTimes(context.Background(), req)
	if err != nil {
		fmt.Println("Error many times", err.Error())
	}

	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error many times", err.Error())
		}

		log.Printf("Recieve %s", msg.Result)
	}
}
