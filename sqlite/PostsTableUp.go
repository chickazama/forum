package sqlite

import "database/sql"

func PostsTableUp(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS "POSTS" (
		PostID INTEGER PRIMARY KEY AUTOINCREMENT,
		Category TEXT NOT NULL,
		Content TEXT NOT NULL,
		Dislikes INTEGER NOT NULL,
		Likes INTEGER NOT NULL,
		CreatedAt BIG INTEGER NOT NULL,
		UpdatedAt BIG INTEGER NOT NULL,
		UserID INTEGER NOT NULL
	);`
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
