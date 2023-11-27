package main

import (
	"context"
	pb "golang-grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("Greet function was invoked with %v\n", req)

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}

func (s *Server) Calc(ctx context.Context, req *pb.CaclRequest) (*pb.CaclResponse, error) {

	log.Printf("Greet function was invoked with %v\n", req)

	return &pb.CaclResponse{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil
}
