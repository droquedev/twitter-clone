package repository

import (
	"database/sql"
	"twitter-clone/server/pkg/post/domain/model"
	"twitter-clone/server/pkg/post/domain/repository"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repository.PostRepository {
	return &postRepository{db}
}

func (r *postRepository) FindAll() ([]*model.Post, error) {
	rows, err := r.db.Query("SELECT * FROM posts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]*model.Post, 0)

	for rows.Next() {
		post := new(model.Post)

		err := rows.Scan(&post.ID, &post.Content, &post.UserID)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *postRepository) FindById(id string) (*model.Post, error) {
	row := r.db.QueryRow("SELECT * FROM posts WHERE id = $1", id)

	post := new(model.Post)

	err := row.Scan(&post.ID, &post.Content, &post.UserID)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) Create(post *model.Post) error {
	_, err := r.db.Exec("INSERT INTO post (id, content, user_id) VALUES ($1, $2, $3)", post.ID, post.Content, post.UserID)

	if err != nil {
		return err
	}

	return nil
}
