package main

import (
	"fmt"
	pb "golang-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	for i := 0; i < 10; i++ {

		stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, number %d", req.FirstName, i),
		})

	}

	return nil
}
