package sqlite

import (
	"database/sql"
	"matthewhope/forum/models"
)

func CreateUser(db *sql.DB, u models.User) (models.User, error) {
	ret := u
	query := `INSERT INTO "USERS" (
		CreatedAt,
		EmailAddress,
		EncryptedPassword,
		FirstName,
		LastName,
		UpdatedAt, 
		Username
	) VALUES (?, ?, ?, ?, ?, ?, ?);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return ret, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.CreatedAt, u.EmailAddress, u.EncryptedPassword, u.FirstName, u.LastName, u.UpdatedAt, u.Username)
	if err != nil {
		return ret, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return ret, err
	}
	ret.UserID = int(id)
	return ret, nil
}
