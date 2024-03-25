package main

import (
	pb "golang-grpc/calculator/proto"
	"io"
	"log"
)

func Avg(stream pb.CalculatorService_AvgServer) error {

	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error hwiele reading client stream %v\n", err)
		}

		sum += req.Number

		count++
	}
}
