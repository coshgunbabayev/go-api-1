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
	ID     string `json:"id"`
	UserID string `json:"-"`
	Text   string `json:"text"`
}

func IsEmpty(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}
