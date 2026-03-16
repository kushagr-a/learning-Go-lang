package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// gin.Context (usually aliased as c in handlers) is a request-scoped
// object essential for handling an HTTP request from start to finish

func NewRouter(database *mongo.Database) *gin.Engine {
	// create gin router
	router := gin.Default() // with logger and recovery

	// health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "OK",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Note API",
			"date":    time.Now().UTC(),
		})
	})

	return router
}
