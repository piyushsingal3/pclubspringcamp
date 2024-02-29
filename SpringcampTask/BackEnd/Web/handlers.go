package handlers

import (
	"context"
	"fmt"
	"mongodbtask/models"
	"mongodbtask/store"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

var validate = validator.New()

func RecentActionsHandler(c *gin.Context, m *store.MongoStore) {
	actions, err := m.QueryRecentActions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return the queried recent actions as JSON
	c.JSON(http.StatusOK, actions)
}
func AddUserHandler(c *gin.Context, m *store.MongoStore) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var userdata models.Users

	if err := c.BindJSON(&userdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(userdata)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	userdata.ID = primitive.NewObjectID()

	result, insertErr := m.UsersCollection.InsertOne(ctx, userdata)
	if insertErr != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)
}

func UsersHandler(c *gin.Context, m *store.MongoStore) {
	userdata, err := m.QueryUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return the queried recent actions as JSON
	c.JSON(http.StatusOK, userdata)
}

// // func GetUser(c *gin.Context, m *store.MongoStore) {
// // 	email := c.Params.ByName("email")
// // 	user, err := m.QueryUsers() //m.QueryUsers(email)
// // 	if err != nil {
// // 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// // 		return
// // 	}

//		// return the queried recent actions as JSON
//		c.JSON(http.StatusOK, user)
//	}
func GetUsersByEmails(c *gin.Context, m *store.MongoStore) {

	email := c.Params.ByName("email")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var users []models.Users

	cursor, err := m.UsersCollection.Find(ctx, bson.M{"email": email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, users)
}
