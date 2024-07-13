package handler

import (
	"go-api-1/models"
	"go-api-1/modules/hash"
	"go-api-1/types"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetUsers(c *gin.Context) {
	var userModel models.UserModel

	users, err := userModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{
		"success": true,
		"users":   users,
	})
}

func GetUser(c *gin.Context) {
	var userModel models.UserModel

	username := c.Param("username")

	user, err := userModel.GetByUsername(username)

	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	if types.IsEmpty(user) {
		c.JSON(404, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}

func CreateUser(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var errors = make(map[string]interface{})

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Name == "" {
		errors["name"] = "name is required"
	}

	if body.Surname == "" {
		errors["surname"] = "surname is required"
	}

	if body.Username == "" {
		errors["username"] = "username is required"
	} else {
		var userModel models.UserModel

		user, _ := userModel.GetByUsername(body.Username)

		if !types.IsEmpty(user) {
			errors["username"] = "username is used"
		}
	}

	if body.Password == "" {
		errors["password"] = "password is required"
	} else if len(body.Password) < 8 {
		errors["password"] = "password must be at least 8 characters long"
	}

	if len(errors) > 0 {
		c.JSON(400, gin.H{
			"success": false,
			"errors":  errors,
		})
		return
	}

	user := types.User{
		ID:       models.GenerateIDForUser(),
		Name:     body.Name,
		Surname:  body.Surname,
		Username: body.Username,
		Password: hash.HashPassword(body.Password),
	}

	var userModel models.UserModel

	userModel.Create(user)

	c.JSON(201, gin.H{
		"success": true,
	})
}

func LoginUser(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var errors = make(map[string]interface{})

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if body.Username == "" {
		errors["username"] = "username is required"
	}

	if body.Password == "" {
		errors["password"] = "password is required"
	}

	if len(errors) > 0 {
		c.JSON(400, gin.H{
			"success": false,
			"errors":  errors,
		})
		return
	}

	var userModel models.UserModel

	user, err := userModel.GetByUsername(body.Username)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if types.IsEmpty(user) {
		errors["username"] = "username is invalid"

		c.JSON(400, gin.H{
			"success": false,
			"errors":  errors,
		})
		return
	}

	if !hash.CheckPasswordHash(body.Password, user.Password) {
		errors["password"] = "password is invalid"

		c.JSON(400, gin.H{
			"success": false,
			"errors":  errors,
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.SetCookie("token", tokenString, int(time.Hour.Seconds()*72), "/", "localhost", false, true)

	c.JSON(201, gin.H{
		"success": true,
	})
}
