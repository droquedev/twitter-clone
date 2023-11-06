package user_handlers

import (
	"net/http"
	user_application "twitter-clone/server/app/user/application"
	user_dto "twitter-clone/server/app/user/domain/dto"
	user_entities "twitter-clone/server/app/user/domain/entities"
	"twitter-clone/server/pkg/password"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUseCase *user_application.UserUseCases
}

func NewUserHandler(userUseCase *user_application.UserUseCases) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var createUserDTO user_dto.UserCreateDTO

	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := password.Hash(createUserDTO.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user := &user_entities.User{
		ID:       uuid.NewString(),
		Username: createUserDTO.Username,
		Email:    createUserDTO.Email,
		Password: password,
	}

	err2 := h.userUseCase.UserCreator(user)

	if err2 != nil {
		if err2.Error() == "username or email already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
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
