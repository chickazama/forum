package api

import "matthewhope/forum/repo"

type UsersHandler struct {
	Repo repo.IRepository
}

func NewUsersHandler(r repo.IRepository) *UsersHandler {
	return &UsersHandler{Repo: r}
}
