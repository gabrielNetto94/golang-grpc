package main

import (
	"context"
	pb "golang-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil
}
