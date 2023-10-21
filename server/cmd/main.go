package main

import (
	"log"
	"net/http"
	"os"
	"twitter-clone/server/api/routes"
	"twitter-clone/server/pkg/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	err := database.InitDatabase(os.Getenv("DB_URI"))

	if err != nil {
		log.Fatal("Error initializing the database: ", err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.InitializeRoutes(r, database.GetDB())
	r.Run(":8080")
}
