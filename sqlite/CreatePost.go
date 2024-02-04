package sqlite

import (
	"database/sql"
	"matthewhope/forum/models"
)

func CreatePost(db *sql.DB, p models.Post) (models.Post, error) {
	ret := p
	query := `INSERT INTO "POSTS" (
		Category,
		Content,
		CreatedAt,
		Dislikes,
		Likes,
		UpdatedAt, 
		UserID
	) VALUES (?, ?, ?, ?, ?, ?, ?);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return ret, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(p.Category, p.Content, p.CreatedAt, p.Dislikes, p.Likes, p.UpdatedAt, p.UserID)
	if err != nil {
		return ret, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return ret, err
	}
	ret.PostID = int(id)
	return ret, nil
}
