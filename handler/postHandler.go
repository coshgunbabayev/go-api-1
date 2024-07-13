package handler

import (
	"go-api-1/models"
	"go-api-1/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPosts

func GetPosts(c *gin.Context) {
	var postModel models.PostModel

	posts, err := postModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{
		"success": true,
		"posts":   posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")

	var postModel models.PostModel

	post, err := postModel.GetByID(id)

	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	if types.IsEmpty(post) {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"success": true,
		"post":    post,
	})
}

func CreatePost(c *gin.Context) {
	var body struct {
		Text string `json:"text"`
	}

	var errors = make(map[string]interface{})

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Text == "" {
		errors["text"] = "text is required"
	}

	if len(errors) > 0 {
		c.JSON(400, gin.H{
			"success": false,
			"errors":  errors,
		})
		return
	}

	var userModel models.UserModel

	post := types.Post{
		ID:     models.GenerateIDForPost(),
		UserID: userModel.GetReqUser(c).ID,
		Text:   body.Text,
	}

	var postModel models.PostModel

	postModel.Create(post)

	c.JSON(201, gin.H{
		"success": true,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var postModel models.PostModel

	post, _ := postModel.GetByID(id)

	if types.IsEmpty(post) {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	var userModel models.UserModel

	if post.UserID != userModel.GetReqUser(c).ID {
		c.JSON(403, gin.H{
			"success": false,
			"message": "You are not authorized to delete this post",
		})
		return
	}

	postModel.DeleteByID(post.ID)

	c.JSON(200, gin.H{
		"success": true,
	})
}
