package handler

import (
	"net/http"
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (h *AuthHandler) CreatePostHandler(jwt string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authLoginDTO dto.AuthLoginDTO

		if err := c.ShouldBindJSON(&authLoginDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()

		if err := validate.Struct(authLoginDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		auth, err := h.authUseCase.Login(authLoginDTO, jwt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User Logged Successfully", "result": auth})
	}
}
