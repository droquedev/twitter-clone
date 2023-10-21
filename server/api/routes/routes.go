package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {
	postsRoutes(router, db)
}
