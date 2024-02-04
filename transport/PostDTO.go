package transport

type PostDTO struct {
	PostID    int
	Category  string
	Content   string
	CreatedAt int64
	Dislikes  int
	Likes     int
	UpdatedAt int64
	UserID    int
}
