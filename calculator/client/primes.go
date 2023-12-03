package main

import (
	"context"
	"fmt"
	pb "golang-grpc/calculator/proto"
	"io"
)

func doPrimes(client pb.CalculatorServiceClient) {

	req := &pb.PrimeRequest{
		Number: 120,
	}

	resp, err := client.Primes(context.Background(), req)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	for {
		res, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error Recv %v", err)
		}

		fmt.Println(res.Result)
	}

}
