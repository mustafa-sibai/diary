package Schema

import (
	"database/sql"
	"diary/Types"
)

func CreateContentTable(database *sql.DB) error {
	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS content (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			body TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}

func AddNewPost(database *sql.DB, title, body string) error {
	_, err := database.Exec(`
		INSERT INTO content (title, body) VALUES (?, ?);
	`, title, body)
	return err
}

func FetchAllPosts(database *sql.DB) ([]Types.Post, error) {
	rows, err := database.Query(`
		SELECT * FROM content ORDER BY created_at DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Types.Post
	for rows.Next() {
		var post Types.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
