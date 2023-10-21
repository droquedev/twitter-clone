package routes

import (
	"database/sql"
	"twitter-clone/server/app/handler"
	"twitter-clone/server/app/repository"
	"twitter-clone/server/app/usecase"

	"github.com/gin-gonic/gin"
)

func postsRoutes(router *gin.Engine, db *sql.DB) {
	productGroup := router.Group("/api/v1/posts")

	postRepository := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postHandler := handler.NewPostHandler(postUsecase)

	{
		productGroup.POST("/", postHandler.CreatePost)
	}
}
