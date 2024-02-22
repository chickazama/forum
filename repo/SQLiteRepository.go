package repo

import (
	"database/sql"
	"log"
	"matthewhope/forum/models"
	"matthewhope/forum/sqlite"
)

const (
	dbDriver       = "sqlite3"
	identityDbPath = "./sqlite/data/Identity.db"
	businessDbPath = "./sqlite/data/Business.db"
)

type SQLiteRepository struct {
	identityDb *sql.DB
	businessDb *sql.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	ret := new(SQLiteRepository)
	db, err := sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	ret.identityDb = db
	db, err = sql.Open(dbDriver, businessDbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	ret.businessDb = db
	err = ret.PostsTableUp()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = ret.UsersTableUp()
	if err != nil {
		log.Fatal(err.Error())
	}
	return ret
}

func (r *SQLiteRepository) PostsTableUp() error {
	return sqlite.PostsTableUp(r.businessDb)
}

func (r *SQLiteRepository) UsersTableUp() error {
	return sqlite.UsersTableUp(r.identityDb)
}

func (r *SQLiteRepository) GetAllUsers() ([]models.User, error) {
	return sqlite.GetAllUsers(r.identityDb)
}

func (r *SQLiteRepository) CreateUser(u models.User) (models.User, error) {
	return sqlite.CreateUser(r.identityDb, u)
}

func (r *SQLiteRepository) GetAllPosts() ([]models.Post, error) {
	return sqlite.GetAllPosts(r.businessDb)
}

func (r *SQLiteRepository) CreatePost(p models.Post) (models.Post, error) {
	return sqlite.CreatePost(r.businessDb, p)
}
