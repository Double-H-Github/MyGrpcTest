package realservices

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-test-01/services"
)

type Server struct {
	productMap map[string] *pb.Product
}

func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func(s *Server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, ok := s.productMap[in.Value]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
	}
	return value, status.New(codes.OK, "").Err()
}