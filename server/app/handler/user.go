package handler

import (
	"net/http"
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUsecase *usecase.UserUseCase
}

func NewUserHandler(userUsecase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var createUserDTO dto.UserCreateDTO

	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.Create(createUserDTO)

	if err != nil {
		if err.Error() == "username or email already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"result":  user,
	})
}
