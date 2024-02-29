package server

import (
	"log"
	handlers "mongodbtask/Web"

	//server "mongodbtask/Web"
	"mongodbtask/store"

	//"net/http"
	//"BackEnd/Web/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Performserver(m *store.MongoStore) {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/Activity/recentActions", func(c *gin.Context) {
		handlers.RecentActionsHandler(c, m)

	})
	router.POST("/user/create", func(c *gin.Context) {
		handlers.AddUserHandler(c, m)

	})
	router.GET("/userData", func(c *gin.Context) {
		handlers.UsersHandler(c, m)

	})
	// router.GET("/usersearch/:email", func(c *gin.Context) {
	// 	handlers.GetUser(c, m)

	// })
	router.GET("/email/:email", func(c *gin.Context) {
		handlers.GetUsersByEmails(c, m)

	})
	// Open connection with MongoDB
	if err := m.OpenConnectionWithMongoDB("mongodb://localhost:27017", "project"); err != nil {
		log.Fatalf("Failed to open connection with MongoDB: %v", err)
	}

	//runs the server with localhost
	if err := router.Run(":9000"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)

	}

}
