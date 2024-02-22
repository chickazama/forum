package repo

import "matthewhope/forum/models"

type IRepository interface {
	PostsTableUp() error
	UsersTableUp() error
	GetAllUsers() ([]models.User, error)
	CreateUser(models.User) (models.User, error)
	GetAllPosts() ([]models.Post, error)
	CreatePost(models.Post) (models.Post, error)
}
