package post_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
