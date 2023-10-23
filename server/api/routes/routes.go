package routes

import (
	"database/sql"
	"twitter-clone/server/config"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB, config *config.Config) {
	authRoutes(router, db, config)
	postsRoutes(router, db, config)
}
