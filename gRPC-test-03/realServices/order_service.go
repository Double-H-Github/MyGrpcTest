package realServices

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-test-03/services"
	"io"
	"log"
	"strings"
	"time"
)

type Server struct {
	orderMap map[string]*pb.Order
}

func (s *Server) AddOrder(ctx context.Context, in *pb.Order) (*pb.OrderID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.orderMap == nil {
		s.orderMap = make(map[string]*pb.Order)
	}
	s.orderMap[in.Id] = in
	return &pb.OrderID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *Server) GetOrder(ctx context.Context, orderID *pb.StringValue) (*pb.Order, error) {

	ord, ok := s.orderMap[orderID.Value]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Order does not exist.", orderID.Value)
	}
	return ord, nil
}

func (s *Server) SearchOrders(searchQuery *pb.StringValue, stream pb.OrderManagement_SearchOrdersServer) (error) {
	for key, order := range s.orderMap{
		log.Print(key, order)
		time.Sleep(300 * time.Millisecond)
		for _ , itemStr := range order.Items{
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				if err := stream.Send(order); err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " +key)
				break
			}
		}
	}
	return fmt.Errorf("This is all ")
}

func (s *Server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	orderStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StringValue{Value:"Orders processed " + orderStr})
		}
		s.orderMap[order.Id] = order

		log.Printf("Order ID ", order.Id, ": Update")
		orderStr += order.Id + ", "
	}
}

func (s *Server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	combinedShipmentMap := make(map[string]*pb.CombinedShipment)
	batchMarker := 0
	orderBatchSize := 3
	for {
		orderId, err := stream.Recv()
		if err == io.EOF {
			for _, comb := range combinedShipmentMap {
				stream.Send(comb)
			}
			return nil
		}
		if err != nil {
			return err
		}

		if ord, ok := s.orderMap[orderId.Value]; !ok {
			continue
		} else {
			if ordList, ok := combinedShipmentMap[ord.Destination]; ok {
				ordList.OrdersList = append(ordList.OrdersList, ord)
			} else {
				tempCombin := &pb.CombinedShipment{
					OrdersList: make([]*pb.Order, 0),
				}
				tempCombin.OrdersList = append(tempCombin.OrdersList, ord)
				combinedShipmentMap[ord.Destination] = tempCombin
			}
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap{
				stream.Send(comb)
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]*pb.CombinedShipment)
		} else {
			batchMarker ++
		}
	}
}