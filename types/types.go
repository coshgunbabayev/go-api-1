package types

import "reflect"

type User struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Post struct {
	ID      int    `json:"id"`
	UserID  string `json:"-"`
	Content string `json:"content"`
}

func IsEmpty(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}
