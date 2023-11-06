package routes

import (
	"database/sql"
	user_application "twitter-clone/server/app/user/application"
	user_handlers "twitter-clone/server/app/user/infrastructure/handlers"
	user_repositories "twitter-clone/server/app/user/infrastructure/repositories"
	"twitter-clone/server/config"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine, db *sql.DB, config *config.Config) {

	userRepository := user_repositories.NewUserPostgresRepository(db)

	userUseCases := user_application.NewUserUseCases(userRepository)

	userHandler := user_handlers.NewUserHandler(userUseCases)

	productGroup := router.Group("/api/v1/users")
	{
		productGroup.POST("/", userHandler.CreateUserHandler)
	}
}
