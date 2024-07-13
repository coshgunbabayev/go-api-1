package models

import (
	"fmt"
	"go-api-1/database"
	"go-api-1/modules/generate"
	"go-api-1/types"
)

type PostModel struct {
}

func GenerateIDForPost() string {
	id := generate.GenerateString(20)

	var postModel PostModel

	post, _ := postModel.GetByID(id)

	if types.IsEmpty(post) {
		return id
	} else {
		return GenerateIDForUser()
	}
}

func (*PostModel) GetAll() ([]types.Post, error) {
	db := database.GetDatabase()
	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []types.Post

	for rows.Next() {
		var post types.Post
		rows.Scan(&post.ID, &post.UserID, &post.Text)
		posts = append(posts, post)
	}

	return posts, nil
}

func (*PostModel) GetByID(id string) (types.Post, error) {
	db := database.GetDatabase()
	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)

	var post types.Post
	row.Scan(&post.ID, &post.UserID, &post.Text)

	return post, nil
}

func (*PostModel) Create(post types.Post) error {
	db := database.GetDatabase()

	_, err := db.Exec("INSERT INTO posts (id, userid, text) VALUES (?, ?, ?)",
		post.ID, post.UserID, post.Text)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (*PostModel) DeleteByID(id string) error {
	db := database.GetDatabase()

	_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
