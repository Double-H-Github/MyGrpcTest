package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test-client/services"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	prodResp, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodResp.ProdStock)
}
