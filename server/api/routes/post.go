package routes

import (
	"database/sql"
	"twitter-clone/server/app/handler"
	"twitter-clone/server/app/repository"
	"twitter-clone/server/app/usecase"

	"github.com/gin-gonic/gin"
)

func postsRoutes(router *gin.Engine, db *sql.DB) {

	postRepository := repository.NewPostPostgresRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postHandler := handler.NewPostHandler(postUsecase)

	productGroup := router.Group("/api/v1/posts")

	{
		productGroup.POST("/", postHandler.CreatePostHandler)
		productGroup.GET("/:user_id", postHandler.FindAllByUserIdHandler)
	}

}
