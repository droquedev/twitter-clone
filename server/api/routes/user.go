package routes

import (
	"database/sql"
	"twitter-clone/server/app/handler"
	"twitter-clone/server/app/repository"
	"twitter-clone/server/app/usecase"
	"twitter-clone/server/config"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine, db *sql.DB, config *config.Config) {

	userRepository := repository.NewUserPostgresRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	productGroup := router.Group("/api/v1/users")
	{
		productGroup.POST("/", userHandler.CreateUserHandler)
	}
}
