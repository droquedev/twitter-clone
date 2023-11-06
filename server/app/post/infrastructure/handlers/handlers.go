package post_handlers

import post_application "twitter-clone/server/app/post/application"

type PostHandler struct {
	postUsecase *post_application.PostUseCase
}

func NewPostHandler(postUsecase *post_application.PostUseCase) *PostHandler {
	return &PostHandler{
		postUsecase: postUsecase,
	}
}
