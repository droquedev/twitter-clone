package handler

import (
	"net/http"
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	postUsecase *usecase.PostUseCase
}

func NewPostHandler(postUsecase *usecase.PostUseCase) *PostHandler {
	return &PostHandler{
		postUsecase: postUsecase,
	}
}

func (h *PostHandler) CreatePostHandler(c *gin.Context) {
	var createPostDTO dto.CreatePostDTO

	if err := c.ShouldBindJSON(&createPostDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(createPostDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postUsecase.Create(createPostDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

func (h *PostHandler) FindAllByUserIdHandler(c *gin.Context) {
	userID := c.Param("user_id")

	posts, err := h.postUsecase.FindAllByUserId(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posts retrieved successfully",
		"result":  posts,
	})
}
