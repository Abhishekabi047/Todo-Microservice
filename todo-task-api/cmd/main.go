package main

import (
	"fmt"
	"log"
	"net"
	"task-svc/pkg/config"
	"task-svc/pkg/db"
	"task-svc/pkg/pb"
	"task-svc/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed at config", err)
	}
	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}

	fmt.Println("Product svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to sever", err)
	}
}
