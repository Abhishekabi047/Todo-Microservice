package routes

import (
	"context"
	"net/http"
	"todo-api/pkg/todo/pb"

	"github.com/gin-gonic/gin"
)

func CompleteList(ctx *gin.Context,c pb.TodoServiceClient){
	res,err:=c.CompleteList(context.Background(),&pb.CompleteListRequest{})

	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway,err)
		return
	}
	ctx.JSON(http.StatusOK,&res)
}