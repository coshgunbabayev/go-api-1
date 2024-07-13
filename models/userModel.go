package models

import (
	"fmt"
	"go-api-1/database"
	"go-api-1/modules/generate"
	"go-api-1/types"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
}

func GenerateIDForUser() string {
	id := generate.GenerateString(20)

	var userModel UserModel

	user, _ := userModel.GetByID(id)

	if types.IsEmpty(user) {
		return id
	} else {
		return GenerateIDForUser()
	}
}

func (*UserModel) GetAll() ([]types.User, error) {
	db := database.GetDatabase()
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User

	for rows.Next() {
		var user types.User
		rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Username, &user.Password)
		users = append(users, user)
	}

	return users, nil
}

func (*UserModel) GetReqUser(c *gin.Context) types.User {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return types.User{}
	}

	return user.(types.User)
}

func (*UserModel) GetByID(id string) (types.User, error) {
	db := database.GetDatabase()
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var user types.User
	row.Scan(&user.ID, &user.Name, &user.Surname, &user.Username, &user.Password)

	return user, nil
}

func (*UserModel) GetByUsername(username string) (types.User, error) {
	db := database.GetDatabase()
	row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)

	var user types.User
	row.Scan(&user.ID, &user.Name, &user.Surname, &user.Username, &user.Password)

	return user, nil
}

func (*UserModel) Create(user types.User) error {
	db := database.GetDatabase()

	_, err := db.Exec("INSERT INTO users (id, name, surname, username, password) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Name, user.Surname, user.Username, user.Password)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
