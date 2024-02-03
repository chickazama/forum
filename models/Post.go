package models

type Post struct {
	PostID    int
	Category  string
	Content   string
	CreatedAt int64
	UpdatedAt int64
	UserID    int
}
