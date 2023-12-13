package main

import (
	"context"
	"fmt"
	pb "golang-grpc/calculator/proto"
	"log"
)

func doAvg(ctx pb.CalculatorServiceClient) {

	stream, err := ctx.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var numbers = []int32{5, 5}

	for _, number := range numbers {
		if err := stream.Send(&pb.AvgRequest{
			Number: number,
		}); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Println(res.Result)

}
