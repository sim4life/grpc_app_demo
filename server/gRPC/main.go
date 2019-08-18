package main

import (
	"fmt"
	"log"
	"net"

	handler "github.com/sim4life/grpc_app_demo/handler/gRPC"
	service "github.com/sim4life/grpc_app_demo/internal/gRPC/service"

	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(7000)
	grpcServer := grpc.NewServer()

	repositoryServiceImpl := handler.NewRepositoryServiceGrpcImpl()
	service.RegisterRepositoryServiceServer(grpcServer, repositoryServiceImpl)

	// start the server
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
