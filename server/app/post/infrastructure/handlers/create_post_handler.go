package post_handlers

import (
	"net/http"
	post_dto "twitter-clone/server/app/post/domain/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *PostHandler) CreatePostHandler(c *gin.Context) {
	var createPostDTO post_dto.CreatePostDTO

	if err := c.ShouldBindJSON(&createPostDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createPostDTO.UserID = c.GetString("sub")

	validate := validator.New()

	if err := validate.Struct(createPostDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.postUsecase.PostCreator(createPostDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"result":  post,
	})
}
