package routes

import (
	"database/sql"
	"twitter-clone/server/api/middleware"
	post_application "twitter-clone/server/app/post/application"
	post_handlers "twitter-clone/server/app/post/infrastructure/handlers"
	post_repositories "twitter-clone/server/app/post/infrastructure/repositories"
	"twitter-clone/server/config"

	"github.com/gin-gonic/gin"
)

func postsRoutes(router *gin.Engine, db *sql.DB, config *config.Config) {

	postRepository := post_repositories.NewPostPostgresRepository(db)
	postUsecase := post_application.NewPostUsecase(postRepository)
	postHandler := post_handlers.NewPostHandler(postUsecase)

	productGroup := router.Group("/api/v1/posts", middleware.AuthMiddleware(config))

	{
		productGroup.POST("/", postHandler.CreatePostHandler)
		productGroup.GET("/:user_id", postHandler.FindAllByUserIdHandler)
	}

}
