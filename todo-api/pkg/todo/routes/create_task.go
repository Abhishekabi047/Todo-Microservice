package routes

import (
	"context"
	"net/http"
	"todo-api/pkg/todo/pb"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequestBody struct {
	Task        string `json:"task"`
	Description string `json:"description"`
	Done        bool   `json:"done" `
}

func CreateTask(ctx *gin.Context, c pb.TodoServiceClient) {
	body := CreateTaskRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.CreateTask(context.Background(), &pb.CreateTaskRequest{
		Task:        body.Task,
		Description: body.Description,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
