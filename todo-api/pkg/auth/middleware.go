package auth

import (
	"context"
	"net/http"
	"strings"
	"todo-api/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

type AuthMiddleConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleWare(svc *ServiceClient) AuthMiddleConfig{
	return AuthMiddleConfig{svc}
}

func(c *AuthMiddleConfig) AuthRequired(ctx *gin.Context) {
	authorization:=ctx.Request.Header.Get("authorization")

	if authorization == ""{
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token:=strings.Split(authorization,"Bearer ")
	if len(token) <2{
	ctx.AbortWithStatus(http.StatusUnauthorized)
	return
	}

	res,err:=c.svc.Client.Validate(context.Background(),&pb.ValidateRequest{
		Token: token[1],
	})

	if err !=nil || res.Status != http.StatusOK{
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("userId",res.UserId)
	ctx.Next()
}
