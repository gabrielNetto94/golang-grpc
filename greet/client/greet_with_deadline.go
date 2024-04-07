package main

import (
	"context"
	"fmt"
	pb "golang-grpc/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {

	log.Println("doGreetWithDeadline was invoked")

	//cria um contexto com timeout específico
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Gabriel",
	}

	res, err := c.GreetWithDeadLine(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpected grpc error: %v\n", e)
			}
		} else {
			log.Fatalf("Error GreetWithDeadLine: %v\n", e)
		}
	}

	fmt.Println("GreetWithDeadLine: ", res.Result)

}
