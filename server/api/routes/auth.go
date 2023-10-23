package routes

import (
	"database/sql"
	"twitter-clone/server/app/handler"
	"twitter-clone/server/app/repository"
	"twitter-clone/server/app/usecase"
	"twitter-clone/server/config"

	"github.com/gin-gonic/gin"
)

func authRoutes(router *gin.Engine, db *sql.DB, config *config.Config) {

	userRepository := repository.NewUserPostgresRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)

	productGroup := router.Group("/api/v1/auth")

	{
		productGroup.POST("/", authHandler.LoginHandler(config.JWT_SECRET))
	}
}
