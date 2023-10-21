package dto

type CreatePostDTO struct {
	Content string `json:"content"`
	UserID  string `json:"userId"`
}
