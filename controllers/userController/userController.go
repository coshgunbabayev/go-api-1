package userController

import (
	"fmt"
	"go-api-1/modules/userModules"
	"go-api-1/types"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := userModules.GetAllUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := userModules.GetUserByID(id)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if types.IsEmpty(user) {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var newUser types.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	fmt.Println("newUser", newUser)

	// createdUser, err := userModules.CreateUser(newUser)

	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(201, createdUser)
}
