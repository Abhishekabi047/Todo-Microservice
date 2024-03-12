package routes

import (
	"context"
	"net/http"
	"todo-api/pkg/todo/pb"

	"github.com/gin-gonic/gin"
)

func ListTasks(ctx *gin.Context,c pb.TodoServiceClient) {
	res,err:=c.ListTask(context.Background(),&pb.ListRequest{})
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway,err)
		return
	}

	ctx.JSON(http.StatusOK,&res)
}