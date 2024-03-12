package routes

import (
	"context"
	"net/http"
	"strconv"
	"todo-api/pkg/todo/pb"

	"github.com/gin-gonic/gin"
)

func DeleteTask(ctx *gin.Context,c pb.TodoServiceClient){
	id,_:=strconv.ParseInt(ctx.Param("id"),10,32)

	res,err:=c.DeleteTask(context.Background(),&pb.DeleteTaskRequest{
		Id: int64(id),
	})

	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway,err)
		return
	}
	ctx.JSON(http.StatusOK,&res)
}