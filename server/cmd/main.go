package main

import (
	"log"
	"net/http"
	"os"
	"twitter-clone/server/pkg/shared/infrastructure/datastore"
	"twitter-clone/server/pkg/shared/infrastructure/router"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	err := datastore.InitDatabase(os.Getenv("DB_URI"))

	if err != nil {
		log.Fatal("Error initializing the database: ", err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.InitializeRoutes(r)
	r.Run(":8080")
}
