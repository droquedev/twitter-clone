package handler

import (
	"log"
	"net/http"
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/entity"
	"twitter-clone/server/app/usecase"
	"twitter-clone/server/pkg/jwt"

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

func (h *AuthHandler) LoginHandler(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authLoginDTO dto.AuthLoginDTO

		if err := c.ShouldBindJSON(&authLoginDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing required fields",
				"error":   err.Error(),
				"result":  nil,
			})
			return
		}

		validate := validator.New()

		if err := validate.Struct(authLoginDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing required fields",
				"error":   err.Error(),
				"result":  nil,
			})
			return
		}

		user, err := h.authUseCase.Login(authLoginDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User or password incorrect",
				"error":   err.Error(),
				"result":  nil,
			})
			return
		}

		token, err := jwt.CreateJWT(user.ID, jwtSecret)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"error":   "Internal Server Error",
				"result":  nil,
			})
			return
		}

		auth := entity.Auth{
			Jwt: token,
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User Logged Successfully",
			"result":  auth,
		})
	}
}
