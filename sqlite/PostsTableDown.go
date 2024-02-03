package sqlite

import "database/sql"

func PostsTableDown(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS "POSTS";`
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
