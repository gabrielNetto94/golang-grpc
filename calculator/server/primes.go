package main

import (
	pb "golang-grpc/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	number := req.Number
	divisor := int64(2)

	for number > 1 {
		if number%int32(divisor) == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= int32(divisor)
		} else {
			divisor++
		}
	}
	return nil
}
