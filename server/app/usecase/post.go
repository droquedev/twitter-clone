package usecase

import (
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/entity"

	"github.com/google/uuid"
)

type PostUseCase struct {
	postRepository entity.PostRepository
}

func NewPostUsecase(postRepo entity.PostRepository) *PostUseCase {
	return &PostUseCase{
		postRepository: postRepo,
	}
}

func (u *PostUseCase) Create(createPostDTO dto.CreatePostDTO) error {

	model := &entity.Post{
		ID:      uuid.NewString(),
		Content: createPostDTO.Content,
		UserID:  createPostDTO.UserID,
	}

	return u.postRepository.Create(model)
}
