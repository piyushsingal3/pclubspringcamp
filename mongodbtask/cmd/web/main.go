package main

import (
	"mongodbtask/store"
	//"mongodbtask/worker"
	//"time"
	"mongodbtask/server"
)

func main() {

	mongoStore := store.NewMongoStore()
	go server.Performserver(mongoStore)

	/*go worker.PerformWork(mongoStore)

	select {
	case <-time.After(time.Hour):

	*/
}

/*package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"mongodbtask/store"
	"net/http"
)

var mongoStore *store.MongoStore

func QueryRecentActionsHandler(c *gin.Context) {
	// Query recent actions from the MongoDB store
	actions, err := mongoStore.QueryRecentActions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the queried recent actions as JSON
	c.JSON(http.StatusOK, actions)
}
func main() {
	//  new MongoStore
	mongoStore = store.NewMongoStore()

	if err := mongoStore.OpenConnectionWithMongoDB("mongodb://localhost:27017", "recentActions"); err != nil {
		log.Fatalf("Failed to open connection with MongoDB: %v", err)
	}

	router := gin.Default()

	router.GET("/recentActions", QueryRecentActionsHandler)

	if err := router.Run(":3400"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}

}
*/
