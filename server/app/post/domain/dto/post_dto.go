package post_dto

type CreatePostDTO struct {
	Content string `json:"content" validate:"required"`
	UserID  string `json:"userId" validate:"required"`
}
