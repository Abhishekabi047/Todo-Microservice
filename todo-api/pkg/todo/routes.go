package todo

import (
	"todo-api/pkg/auth"
	"todo-api/pkg/todo/routes"
	"todo-api/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,c *config.Config,authsvc *auth.ServiceClient) {
	a:=auth.InitAuthMiddleWare(authsvc)

	svc:=&ServiceClient{
		Client: InitServiceClient(c),
	}

	routes:=r.Group("/task")
	routes.Use(a.AuthRequired)
	routes.POST("/",svc.CreateTask)
	routes.GET("/",svc.ListTasks)
	routes.DELETE("/:id",svc.DeleteTask)
	routes.GET("/complete",svc.CompleteList)
	routes.POST("/complete/:id",svc.Complete)

}

func(svc *ServiceClient) CreateTask(ctx *gin.Context) {
	routes.CreateTask(ctx,svc.Client)
}

func(svc *ServiceClient) ListTasks(ctx *gin.Context) {
	routes.ListTasks(ctx,svc.Client)
}

func(svc *ServiceClient) DeleteTask(ctx *gin.Context) {
	routes.DeleteTask(ctx,svc.Client)
}

func(svc *ServiceClient) CompleteList(ctx *gin.Context) {
	routes.CompleteList(ctx,svc.Client)
}

func(svc *ServiceClient) Complete(ctx *gin.Context) {
	routes.Complete(ctx,svc.Client)
}