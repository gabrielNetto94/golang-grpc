package main

import (
	"context"
	"fmt"
	pb "golang-grpc/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {

	resp, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	fmt.Println("Result :", resp.Result)
}
