package main

import (
	"log"
	"todo-api/pkg/auth"
	"todo-api/pkg/config"
	"todo-api/pkg/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	c,err:=config.LoadConfig()
	if err != nil{
		log.Fatalln("failed to config",err)
		return
	}
	r:=gin.Default()

	authsvc:=*auth.RegisterRoutes(r,&c)
	todo.RegisterRoutes(r,&c,&authsvc)
	r.Run(c.Port)
}