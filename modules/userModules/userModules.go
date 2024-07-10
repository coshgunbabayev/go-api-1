package userModules

import (
	"encoding/json"
	"go-api-1/types"
	"io/ioutil"
	"os"
)

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
