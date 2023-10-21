package repository

import (
	"database/sql"
	"twitter-clone/server/app/entity"
)

type PostPostgresRepository struct {
	db *sql.DB
}

func NewPostPostgresRepository(db *sql.DB) entity.PostRepository {
	return &PostPostgresRepository{db}
}

func (r *PostPostgresRepository) FindAll() ([]*entity.Post, error) {
	rows, err := r.db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]*entity.Post, 0)

	for rows.Next() {
		post := new(entity.Post)

		err := rows.Scan(&post.ID, &post.Content, &post.UserID)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostPostgresRepository) FindById(id string) (*entity.Post, error) {
	row := r.db.QueryRow("SELECT * FROM posts WHERE id = $1", id)

	post := new(entity.Post)

	err := row.Scan(&post.ID, &post.Content, &post.UserID)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostPostgresRepository) Create(post *entity.Post) error {
	_, err := r.db.Exec("INSERT INTO posts (id, content, user_id) VALUES ($1, $2, $3)", post.ID, post.Content, post.UserID)

	if err != nil {
		return err
	}

	return nil
}
