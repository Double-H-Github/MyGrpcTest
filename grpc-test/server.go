package main

import (
	"google.golang.org/grpc"
	"grpc_test/services"
	"net/http"
)

func main()  {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	//lis, _ := net.Listen("tcp", ":8081")
	//
	//rpcServer.Serve(lis)



	//http interface
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rpcServer.ServeHTTP(writer, request)
	})
	httpServer := &http.Server{
		Addr:              ":8082",
		Handler:           mux,
	}
	httpServer.ListenAndServe()
}