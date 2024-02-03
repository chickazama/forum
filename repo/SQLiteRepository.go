package repo

import (
	"database/sql"
	"log"
)

const (
	dbDriver       = "sqlite3"
	identityDbPath = "./sqlite/data/identity.db"
	businessDbPath = "./sqlite/data/business.db"
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
	return ret
}