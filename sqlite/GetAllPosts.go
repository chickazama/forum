package sqlite

import (
	"database/sql"
	"matthewhope/forum/models"
)

func GetAllPosts(db *sql.DB) ([]models.Post, error) {
	var ret []models.Post
	query := `SELECT * FROM "POSTS";`
	rows, err := db.Query(query)
	if err != nil {
		return ret, err
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.PostID, &p.Category, &p.Content, &p.CreatedAt, &p.Dislikes, &p.Likes, &p.UpdatedAt, &p.UserID)
		if err != nil {
			return ret, err
		}
		ret = append(ret, p)
	}
	return ret, nil
}
