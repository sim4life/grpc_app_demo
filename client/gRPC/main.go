package main

import (
	"context"
	"fmt"

	"github.com/sim4life/grpc_app_demo/internal/gRPC/domain"

	"github.com/sim4life/grpc_app_demo/internal/gRPC/service"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewRepositoryServiceClient(conn)

	for i := range [10]int{} {
		repositoryModel := domain.Repository{
			Id:        int64(i),
			IsPrivate: true,
			Name:      string("Grpc-Demo"),
			UserId:    1245,
		}

		if responseMessage, e := client.Add(context.Background(), &repositoryModel); e != nil {
			panic(fmt.Sprintf("Was not able to insert Record %v", e))
		} else {
			fmt.Println("Record Inserted..")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
