package handlers

import (
	"go-api-1/models"

	"github.com/gin-gonic/gin"
)

func GetHomePage(c *gin.Context) {
	var userModel models.UserModel

	c.HTML(200, "index.html", gin.H{
		"user": userModel.GetReqUser(c),
	})
}

func GetAccountPage(c *gin.Context) {
	c.HTML(200, "account.html", nil)
}

func GetUserPage(c *gin.Context) {
	username := c.Param("username")
	var userModel models.UserModel

	c.HTML(200, "user.html", gin.H{
		"user":     userModel.GetReqUser(c),
		"username": username,
	})
}

func GetPostPage(c *gin.Context) {
	id := c.Param("id")
	var userModel models.UserModel

	c.HTML(200, "post.html", gin.H{
		"user": userModel.GetReqUser(c),
		"id":   id,
	})
}
