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

func (u *PostUseCase) Create(createPostDTO dto.CreatePostDTO) (*entity.Post, error) {

	post := &entity.Post{
		ID:      uuid.NewString(),
		Content: createPostDTO.Content,
		UserID:  createPostDTO.UserID,
	}

	err := u.postRepository.Create(post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (u *PostUseCase) FindAllByUserId(userID string) ([]*entity.Post, error) {
	return u.postRepository.FindAllByUserId(userID)
}
