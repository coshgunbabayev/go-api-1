package main

import (
	"go-api-1/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.GET("/", controller.GetHomePage)
	r.GET("/account", controller.GetAccountPage)

	api := r.Group("/api")
	{
		userAPI := api.Group("/user")
		{
			userAPI.GET("/", controller.GetUsers)
			userAPI.GET("/:username", controller.GetUser)
			userAPI.POST("/signup", controller.CreateUser)
			userAPI.POST("/login", controller.LoginUser)
		}

		// postAPI := api.Group("/post")
		// {
		// 	postAPI.GET("/", controller.GetPosts)
		//     postAPI.GET("/:id", controller.GetPost)
		//     postAPI.POST("/", controller.CreatePost)
		//     postAPI.PUT("/:id", controller.UpdatePost)
		//     postAPI.DELETE("/:id", controller.DeletePost)
		// }
	}

	// r.GET("/api/salam/:aaa", middleware.Authenticate(), func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"message": "Salam " + ctx.Param("aaa")})
	// })

	err := r.Run()
	if err != nil {
		panic("Failed to start the server")
	}
}
