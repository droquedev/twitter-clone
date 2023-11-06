package routes

import (
	"database/sql"
	auth_application "twitter-clone/server/app/auth/application"
	auth_handlers "twitter-clone/server/app/auth/infrastructure/handlers"
	user_repositories "twitter-clone/server/app/user/infrastructure/repositories"
	"twitter-clone/server/config"

	"github.com/gin-gonic/gin"
)

func authRoutes(router *gin.Engine, db *sql.DB, config *config.Config) {

	userRepository := user_repositories.NewUserPostgresRepository(db)
	authUsecase := auth_application.NewAuthLogin(userRepository)
	authHandler := auth_handlers.NewAuthHandler(authUsecase)

	productGroup := router.Group("/api/v1/auth")

	{
		productGroup.POST("/", authHandler.LoginHandler(config.JWT_SECRET))
	}
}
