package auth

import (
	"fmt"
	"todo-api/pkg/auth/pb"
	"todo-api/pkg/config"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("cound nto connect", err)
	}
	return pb.NewAuthServiceClient(cc)
}
