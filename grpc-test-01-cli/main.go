package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-test-01-cli/services"
	"log"
	"time"
)

const address = "localhost:50051"

const address3 = "localhost:50053"

func main()  {

	//test1
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := pb.NewProductInfoClient(conn)
	//
	//name := "Apple iPhone 11"
	//description := "This is a iphone 11"
	//price := float32(1000.0)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.AddProduct(ctx, &pb.Product{
	//	Name:        name,
	//	Description: description,
	//	Price:       price,
	//})
	//if err != nil {
	//	log.Fatalf("Could not add product: %v", err)
	//}
	//log.Printf("Product ID: %s added successfully", r.Value)
	//product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	//if err != nil {
	//	log.Fatalf("Could not get product: %v", err)
	//}
	//log.Printf("Product: ", product.String())

	//test3
	conn, err := grpc.Dial(address3, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4 * time.Second)
	defer cancel()

	orderID, err := c.AddOrder(ctx, &pb.Order{
		Items:       []string{"qwe","asd","zxc"},
		Description: "a",
		Price:       1,
		Destination: "China",
	})
	if err != nil {
		log.Fatalf("Could not add order: %v", err)
	}
	log.Printf("Order ID: %s added successfully", orderID.Value)

	orderID2, err2 := c.AddOrder(ctx, &pb.Order{
		Items:       []string{"qwe"},
		Description: "b",
		Price:       2,
		Destination: "Japan",
	})
	if err2 != nil {
		log.Fatalf("Could not add order: %v", err)
	}
	log.Printf("Order ID: %s added successfully", orderID2.Value)

	orderID3, err3 := c.AddOrder(ctx, &pb.Order{
		Items:       []string{"qwe"},
		Description: "b",
		Price:       2,
		Destination: "Japan",
	})
	if err3 != nil {
		log.Fatalf("Could not add order: %v", err3)
	}
	log.Printf("Order ID: %s added successfully", orderID3.Value)


	//product, err := c.GetOrder(ctx, &pb.StringValue{Value: orderID2.Value})
	//if err != nil {
	//	log.Fatalf("Could not get product: %v", err)
	//}
	//log.Printf("Product: ", product.String())
	//
	//searchStream, _ := c.SearchOrders(ctx, &pb.StringValue{Value: "qwe"})
	//
	//for {
	//	searchOrder, err := searchStream.Recv()
	//	if err != nil {
	//		log.Fatalf("searchOrder error : %v", err)
	//		break
	//	}
	//	log.Print("Search Result : ", searchOrder)
	//
	//}
	updateStream, err := c.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("v.UpdateOrders(_) = _, %v", c, err)
	}

	updateOrder1 := &pb.Order{
		Id: orderID.Value,
		Items:       []string{"qwe"},
		Description: "b",
		Price:       2,
		Destination: "Japan",
	}

	if err := updateStream.Send(updateOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updateOrder1, err)
	}

	updateOrder2 := &pb.Order{
		Id: orderID2.Value,
		Items:       []string{"qwe"},
		Description: "b",
		Price:       2,
		Destination: "Japan",
	}

	if err := updateStream.Send(updateOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updateOrder2, err)
	}

	updateOrder3 := &pb.Order{
		Id: orderID3.Value,
		Items:       []string{"qwe"},
		Description: "b",
		Price:       2,
		Destination: "Japan",
	}

	if err := updateStream.Send(updateOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updateOrder3, err)
	}

	updatesResp, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}
	log.Printf("Update Orders Res : %s", updatesResp)

	searchStream, _ := c.SearchOrders(ctx, &pb.StringValue{Value: "qwe"})

	for {
		searchOrder, err := searchStream.Recv()
		if err != nil {
			log.Fatalf("searchOrder error : %v", err)
			break
		}
		log.Print("Search Result : ", searchOrder)

	}


}