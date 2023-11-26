package main

import (
	"context"
	"fmt"
	pb "golang-grpc/greet/proto"
)

func doGreet(ctx pb.GreetServiceClient) {

	fmt.Println("Greet client was invoked")

	resp, err := ctx.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Gabriel Netto",
	})
	if err != nil {
		fmt.Println("Error request:", err.Error())
	}
	fmt.Println("Response: ", resp.Result)
}
