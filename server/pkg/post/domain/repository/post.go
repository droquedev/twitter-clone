package repository

import "twitter-clone/server/pkg/post/domain/model"

type PostRepository interface {
	FindAll() ([]*model.Post, error)
	FindById(id string) (*model.Post, error)
	Create(post *model.Post) error
}
