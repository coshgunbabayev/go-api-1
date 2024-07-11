package userController

import (
	"go-api-1/modules/userModules"
	"go-api-1/types"
	"net/http"

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
	username := c.Param("username")

	user, err := userModules.GetUserByUsername(username)

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
	var newUser struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var errors = make(map[string]interface{})

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Name == "" {
		errors["name"] = "name is required"
	}

	if newUser.Surname == "" {
		errors["surname"] = "surname is required"
	}

	if newUser.Username == "" {
		errors["username"] = "username is required"
	} else {
		user, _ := userModules.GetUserByUsername(newUser.Username)

		if !types.IsEmpty(user) {
			errors["username"] = "username id used"
		}
	}

	if newUser.Password == "" {
		errors["password"] = "password is required"
	} else if len(newUser.Password) < 8 {
		errors["password"] = "password must be at least 8 characters long"
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errors":  errors,
		})
		return
	}

	user := types.User{
		ID:       userModules.GenerateID(),
		Name:     newUser.Name,
		Surname:  newUser.Surname,
		Username: newUser.Username,
		Password: newUser.Password,
	}

	userModules.AddUser(user)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"success": true,
	})
}
