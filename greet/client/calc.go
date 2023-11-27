package main

import (
	"context"
	"fmt"
	pb "golang-grpc/greet/proto"
)

func calc(ctx pb.GreetServiceClient) {
	fmt.Println("Calc client was invoked")

	resp, err := ctx.Calc(context.Background(), &pb.CaclRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		fmt.Println("Error request:", err.Error())
	}
	fmt.Println("Response: ", resp.Result)
}
