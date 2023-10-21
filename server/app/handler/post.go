package handler

import (
	"net/http"
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/usecase"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUsecase *usecase.PostUseCase
}

func NewPostHandler(postUsecase *usecase.PostUseCase) *PostHandler {
	return &PostHandler{
		postUsecase: postUsecase,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var createPostDTO dto.CreatePostDTO

	if err := c.ShouldBindJSON(&createPostDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postUsecase.Create(createPostDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
