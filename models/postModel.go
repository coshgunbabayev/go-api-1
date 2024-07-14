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

	if post.IsEmpty() {
		return id
	} else {
		return GenerateIDForUser()
	}
}

func (*PostModel) GetAll() ([]types.Post, error) {
	db := database.GetDatabase()
	rows, err := db.Query("SELECT * FROM posts WHERE topostid = ?", "")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []types.Post

	for rows.Next() {
		var post types.Post
		rows.Scan(&post.ID, &post.UserID, &post.ToPostID, &post.Text)
		posts = append(posts, post)
	}

	for i, post := range posts {
		var userModel UserModel
		var postModel PostModel
		user, _ := userModel.GetByID(post.UserID)
		posts[i].User = user
		likes, _ := postModel.GetLikesByID(post.ID)
		posts[i].Likes = likes
		comments, _ := postModel.GetCommentsById(post.ID)
		posts[i].Comments = comments
	}

	return posts, nil
}

func (*PostModel) GetByID(id string) (types.Post, error) {
	db := database.GetDatabase()
	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)

	var post types.Post
	row.Scan(&post.ID, &post.UserID, &post.ToPostID, &post.Text)

	var userModel UserModel
	var postModel PostModel
	user, _ := userModel.GetByID(post.UserID)
	post.User = user
	likes, _ := postModel.GetLikesByID(post.ID)
	post.Likes = likes
	comments, _ := postModel.GetCommentsById(post.ID)
	post.Comments = comments

	return post, nil
}

func (*PostModel) GetLikesByID(id string) ([]types.User, error) {
	db := database.GetDatabase()
	rows, err := db.Query("SELECT userid FROM likes WHERE postid =?", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	var userModel UserModel

	for rows.Next() {
		var user types.User
		rows.Scan(&user.ID)
		user, _ = userModel.GetByID(user.ID)
		users = append(users, user)
	}

	return users, nil
}

func (*PostModel) GetCommentsById(id string) ([]types.Post, error) {
	db := database.GetDatabase()
	rows, err := db.Query("SELECT * FROM posts WHERE topostid =?", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []types.Post

	for rows.Next() {
		var post types.Post
		rows.Scan(&post.ID, &post.UserID, &post.ToPostID, &post.Text)
		posts = append(posts, post)
	}

	for i, post := range posts {
		var userModel UserModel
		var postModel PostModel
		user, _ := userModel.GetByID(post.UserID)
		posts[i].User = user
		likes, _ := postModel.GetLikesByID(post.ID)
		posts[i].Likes = likes
		comments, _ := postModel.GetCommentsById(post.ID)
		posts[i].Comments = comments
	}

	return posts, nil
}

func (*PostModel) IsLikedByID(userId string, postId string) bool {
	db := database.GetDatabase()
	row := db.QueryRow("SELECT COUNT(*) FROM likes WHERE userid =? AND postid =?", userId, postId)

	var count int
	row.Scan(&count)

	return count > 0
}

func (*PostModel) LikeByID(userId string, postId string) {
	db := database.GetDatabase()
	db.Exec("INSERT INTO likes (userid, postid) VALUES (?,?)", userId, postId)
}

func (*PostModel) UnlikeByID(userId string, postId string) {
	db := database.GetDatabase()
	db.Exec("DELETE FROM likes WHERE userid =? AND postid =?", userId, postId)
}

func (*PostModel) CreateAsPost(post types.Post) error {
	db := database.GetDatabase()

	_, err := db.Exec("INSERT INTO posts (id, userid, topostid, text) VALUES (?, ?, ?, ?)",
		post.ID, post.UserID, "", post.Text)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (*PostModel) CreateAsComment(post types.Post) error {
	db := database.GetDatabase()

	_, err := db.Exec("INSERT INTO posts (id, userid, topostid, text) VALUES (?,?,?,?)",
		post.ID, post.UserID, post.ToPostID, post.Text)

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
