package types

type User struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Post struct {
	ID       string `json:"id"`
	UserID   string `json:"-"`
	User     User   `json:"user"`
	ToPostID string `json:"topostid"`
	Text     string `json:"text"`
	Likes    []User `json:"likes"`
	Comments []Post `json:"comments"`
}

func (user User) IsEmpty() bool {
	return user.ID == ""
}

func (post Post) IsEmpty() bool {
	return post.ID == ""
}
