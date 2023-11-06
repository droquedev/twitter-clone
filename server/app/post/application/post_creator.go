package post_application

import (
	post_dto "twitter-clone/server/app/post/domain/dto"
	post_entities "twitter-clone/server/app/post/domain/entities"

	"github.com/google/uuid"
)

func (u *PostUseCase) PostCreator(createPostDTO post_dto.CreatePostDTO) (*post_entities.Post, error) {

	post := &post_entities.Post{
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

func (u *PostUseCase) FindAllByUserId(userID string) ([]*post_entities.Post, error) {
	return u.postRepository.FindAllByUserId(userID)
}
