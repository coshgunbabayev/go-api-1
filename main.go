package main

import (
	"go-api-1/handler"
	"go-api-1/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	page := r.Group("/")
	{
		page.GET("/", middleware.AuthenticateForPage(), handler.GetHomePage)
		page.GET("/account", handler.GetAccountPage)
		page.GET("/user/:username", middleware.AuthenticateForPage(), handler.GetUserPage)
	}

	api := r.Group("/api")
	{
		userAPI := api.Group("/user")
		{
			userAPI.GET("/", handler.GetUsers)
			userAPI.GET("/:username", middleware.AuthenticateForAPI(), handler.GetUser)
			userAPI.POST("/signup", handler.CreateUser)
			userAPI.POST("/login", handler.LoginUser)
		}

		postAPI := api.Group("/post")
		{
			postAPI.GET("/", middleware.AuthenticateForAPI(), handler.GetPosts)
			postAPI.GET("/:id", middleware.AuthenticateForAPI(), handler.GetPost)
			postAPI.POST("/", middleware.AuthenticateForAPI(), handler.CreatePost)
			postAPI.DELETE("/:id", middleware.AuthenticateForAPI(), handler.DeletePost)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		panic("Failed to start the server")
	}
}
