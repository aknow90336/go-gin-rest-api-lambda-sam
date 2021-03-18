package main

import (
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	. "todolist/src"
)

func main() {
	// Listen port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := ":" + port

	router := gin.New()

	// Set routing
	routerGroup := router.Group("/api")
	{
		routerGroup.GET("/todo", Wrapper(GetTodoItems))
		routerGroup.POST("/todo", Wrapper(CreateTodoItem))
		routerGroup.POST("/todo/:id", Wrapper(UpdateTodoItem))
		routerGroup.DELETE("/todo/:id", Wrapper(DeleteTodoItem))
	}

	log.Fatal(gateway.ListenAndServe(addr, router))
}
