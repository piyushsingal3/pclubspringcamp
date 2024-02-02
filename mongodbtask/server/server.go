package server

import (
	"log"
	"mongodbtask/store"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Performserver(m *store.MongoStore, wg *sync.WaitGroup) {
	defer wg.Done()
	// Creates a new gin router
	router := gin.Default()

	// Define a route "/recentActions"
	router.GET("/recentActions", func(c *gin.Context) {

		actions, err := m.QueryRecentActions()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the queried recent actions as JSON
		c.JSON(http.StatusOK, actions)
	})

	// Open connection with MongoDB
	if err := m.OpenConnectionWithMongoDB("mongodb://localhost:27017", "recent"); err != nil {
		log.Fatalf("Failed to open connection with MongoDB: %v", err)
	}

	//runs the server with localhost
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}

}
