package main

import (
	"google.golang.org/grpc"
	"grpc-test-01/realservices"
	pb "grpc-test-01/services"
	"log"
	"net"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &realservices.Server{})

	log.Printf("Starting grpc listen on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}