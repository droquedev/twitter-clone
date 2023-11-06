package post_application

import post_entities "twitter-clone/server/app/post/domain/entities"

type PostUseCase struct {
	postRepository post_entities.PostRepository
}

func NewPostUsecase(postRepo post_entities.PostRepository) *PostUseCase {
	return &PostUseCase{
		postRepository: postRepo,
	}
}
