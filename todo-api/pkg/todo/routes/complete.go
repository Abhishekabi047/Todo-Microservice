package routes

import (
	"context"
	"net/http"
	"strconv"
	"todo-api/pkg/todo/pb"

	"github.com/gin-gonic/gin"
)

func Complete(ctx *gin.Context,c pb.TodoServiceClient){
	id,_:=strconv.ParseInt(ctx.Param("id"),10,32)

	res,err:=c.Complete(context.Background(),&pb.CompleteRequest{
		Id: int64(id),
	})

	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway,err)
		return
	}
	ctx.JSON(http.StatusOK,&res)
}