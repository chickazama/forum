package sqlite

import "database/sql"

func UsersTableDown(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS "USERS";`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
