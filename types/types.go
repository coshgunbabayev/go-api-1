package types

import "reflect"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	ID      int    `json:"id"`
	UserID  string `json:"userid"`
	Content string `json:"content"`
}

func IsEmpty(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}
