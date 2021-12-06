package main

import (
	"google.golang.org/grpc"
	"grpc-test-03/realServices"
	pb "grpc-test-03/services"
	"log"
	"net"
)

const port = ":50053"

func main() {
	//绑定tcp端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//注册grpc服务器
	s := grpc.NewServer()
	//向grpc服务器上注册服务
	pb.RegisterOrderManagementServer(s, &realServices.Server{})
	log.Printf("Starting grpc listen on port " + port)

	//将服务器与端口号绑定
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}