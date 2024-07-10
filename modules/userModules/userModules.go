package userModules

import (
	"encoding/json"
	"fmt"
	"go-api-1/modules/generate"
	"go-api-1/types"
	"io/ioutil"
	"os"
)

func GenerateID() string {
	id := generate.GenerateString(10)

	user, _ := GetUserByID(id)

	if types.IsEmpty(user) {
		return id
	} else {
		return GenerateID()
	}
}

func GetAllUsers() ([]types.User, error) {
	file, err := os.Open("./db/users.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var users []types.User
	json.Unmarshal(byteValue, &users)

	return users, nil
}

func GetUserByID(id string) (types.User, error) {
	users, err := GetAllUsers()
	if err != nil {
		return types.User{}, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return types.User{}, nil
}

func GetUserByUsername(username string) (types.User, error) {
	users, err := GetAllUsers()
	if err != nil {
		return types.User{}, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return types.User{}, nil
}

func UsernameIsUsed(username string) bool {
	users, _ := GetAllUsers()

	for _, user := range users {
		if user.Username == username {
			return true
		}
	}

	return false
}

func AddUser(user types.User) {
	users, _ := GetAllUsers()

	users = append(users, user)

	file, err := os.Create("./db/users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Write(jsonData)
}
