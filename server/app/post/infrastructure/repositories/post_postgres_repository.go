package post_repositories

import (
	"database/sql"
	post_entities "twitter-clone/server/app/post/domain/entities"
)

type PostPostgresRepository struct {
	db *sql.DB
}

func NewPostPostgresRepository(db *sql.DB) post_entities.PostRepository {
	return &PostPostgresRepository{db}
}

func (r *PostPostgresRepository) Create(post *post_entities.Post) error {
	_, err := r.db.Exec("INSERT INTO posts (id, content, user_id) VALUES ($1, $2, $3)", post.ID, post.Content, post.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostPostgresRepository) FindAllByUserId(userID string) ([]*post_entities.Post, error) {
	rows, err := r.db.Query("SELECT * FROM posts WHERE user_id = $1", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]*post_entities.Post, 0)

	for rows.Next() {
		post := new(post_entities.Post)

		err := rows.Scan(&post.ID, &post.Content, &post.UserID)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostPostgresRepository) FindById(id string) (*post_entities.Post, error) {
	row := r.db.QueryRow("SELECT * FROM posts WHERE id = $1", id)

	post := new(post_entities.Post)

	err := row.Scan(&post.ID, &post.Content, &post.UserID)

	if err != nil {
		return nil, err
	}

	return post, nil
}
