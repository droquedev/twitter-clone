package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"twitter-clone/server/pkg/post/domain/model"
	"twitter-clone/server/pkg/post/infrastructure/repository"
	"twitter-clone/server/pkg/shared/infrastructure/datastore"

	"github.com/gin-gonic/gin"
)

type postDTO struct {
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

func CreatePost(c *gin.Context) {
	var postBody postDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&postBody); err != nil {
		fmt.Printf("error %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request decode",
			"error":   err,
		})
		return
	}

	var post = model.Post{
		Content: postBody.Content,
		UserID:  postBody.UserID,
		ID:      "48013754-9dbb-4991-b352-f71b91713e05",
	}

	repository := repository.NewPostRepository(datastore.GetDB())

	if err := repository.Create(&post); err != nil {
		fmt.Printf("error %s", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post created",
		"X":       post,
	})
}
