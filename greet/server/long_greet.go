package main

import (
	"fmt"
	pb "golang-grpc/greet/proto"
	"io"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	res := ""
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			fmt.Println("Error msg:", err.Error())
			break
		}

		res += "Hello " + req.FirstName + "\n"
	}

	return nil
}
