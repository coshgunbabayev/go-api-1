package main

import (
	"go-api-1/controllers/pageController"
	"go-api-1/controllers/userController"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.GET("/", pageController.GetHomePage)
	r.GET("/signup", pageController.GetSignUpPage)
	r.GET("/login", pageController.GetLoginPage)

	r.GET("/api/user", userController.GetUsers)
	r.GET("/api/user/:username", userController.GetUser)
	r.POST("/api/user/signup", userController.CreateUser)

	err := r.Run()
	if err != nil {
		panic("Failed to start the server")
	}
}
