package main

import (
	"go-api-1/handlers"
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
		page.GET("/", middleware.AuthenticateForPage(), handlers.GetHomePage)
		page.GET("/account", handlers.GetAccountPage)
		page.GET("/user/:username", middleware.AuthenticateForPage(), handlers.GetUserPage)
		page.GET("/post/:id", middleware.AuthenticateForPage(), handlers.GetPostPage)
	}

	api := r.Group("/api")
	{
		userAPI := api.Group("/user")
		{
			userAPI.GET("/", handlers.GetUsers)
			userAPI.GET("/:username", middleware.AuthenticateForAPI(), handlers.GetUser)
			userAPI.POST("/signup", handlers.CreateUser)
			userAPI.POST("/login", handlers.LoginUser)
		}

		postAPI := api.Group("/post")
		{
			postAPI.POST("/", middleware.AuthenticateForAPI(), handlers.CreatePost)
			postAPI.GET("/", middleware.AuthenticateForAPI(), handlers.GetPosts)
			postAPI.GET("/:id", middleware.AuthenticateForAPI(), handlers.GetPost)
			postAPI.GET("/:id/like", middleware.AuthenticateForAPI(), handlers.GetLikeCaseOfPost)
			postAPI.POST("/:id", middleware.AuthenticateForAPI(), handlers.CreateComment)
			postAPI.DELETE("/:id", middleware.AuthenticateForAPI(), handlers.DeletePost)
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
