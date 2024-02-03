package repo

import "matthewhope/forum/models"

type IRepository interface {
	GetAllUsers() ([]models.User, error)
}
