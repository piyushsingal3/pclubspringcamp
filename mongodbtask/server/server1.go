package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"mongodbtask/store"
	"net/http"
)

//var mongoStore *store.MongoStore

func Performserver(m *store.MongoStore) {
	//  new MongoStore
	var c *gin.Context
	actions, err := m.QueryRecentActions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the queried recent actions as JSON
	c.JSON(http.StatusOK, actions)

	if err := m.OpenConnectionWithMongoDB("mongodb://localhost:27017", "recentActions"); err != nil {
		log.Fatalf("Failed to open connection with MongoDB: %v", err)
	}

	// Create a Gin router
	router := gin.Default()

	router.GET("/recentActions", c.Handler())

	// Run the server
	if err := router.Run(":8490"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
