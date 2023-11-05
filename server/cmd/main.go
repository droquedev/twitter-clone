package main

import (
	"log"
	"net/http"
	"twitter-clone/server/api/middleware"
	"twitter-clone/server/api/routes"
	"twitter-clone/server/app/event"
	"twitter-clone/server/config"
	"twitter-clone/server/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	config, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	event.InitSubscriptions()
	db, err := database.InitDatabase(config)

	r.Use(middleware.GlobalErrorHandler())

	if err != nil {
		log.Fatal("Error initializing the database: ", err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.InitializeRoutes(r, db, config)

	r.Run(":8080")
}
