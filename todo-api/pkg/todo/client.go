package todo

import (
	"fmt"
	"todo-api/pkg/config"
	"todo-api/pkg/todo/pb"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.TodoServiceClient
}

func InitServiceClient(c *config.Config) pb.TodoServiceClient {
	cc, err := grpc.Dial(c.TodoSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to conect", err)
	}
	return pb.NewTodoServiceClient(cc)
}
