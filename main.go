package main

import (
	"go-api-1/controllers/userController"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/user", userController.GetUsers)
	r.GET("/api/user/:id", userController.GetUser)
	r.POST("/api/user/signup", userController.CreateUser)

	err := r.Run()
	if err != nil {
		panic("Failed to start the server")
	}
}
