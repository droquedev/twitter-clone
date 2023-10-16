package router

import (
	"twitter-clone/server/pkg/post/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func postsRoutes(router *gin.Engine) {
	productGroup := router.Group("/api/v1/posts")
	{
		productGroup.POST("/", controller.CreatePost)
	}
}
