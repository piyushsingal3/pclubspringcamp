package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mongodbtask/models"
)

type MongoStore struct {
	RecentActionsCollection *mongo.Collection
	UsersCollection         *mongo.Collection
}

func NewMongoStore() *MongoStore {
	return &MongoStore{}
}

func (m *MongoStore) OpenConnectionWithMongoDB(connectionString, databaseName string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	db := client.Database(databaseName)

	m.RecentActionsCollection = db.Collection("recentActions")
	m.UsersCollection = db.Collection("Users")

	return nil
}

func (m *MongoStore) StoreRecentActionsInTheDatabase(actions []models.RecentActions) error {

	documents := make([]interface{}, len(actions))
	for i, action := range actions {

		documents[i] = action
	}

	_, err := m.RecentActionsCollection.InsertMany(context.TODO(), documents)
	if err != nil {
		return fmt.Errorf("failed to store recentActions in the database: %v", err)
	}

	return nil
}
func (m *MongoStore) StoreUsersInTheDatabase(userdata []models.Users) error {

	documents := make([]interface{}, len(userdata))
	for i, userdata := range userdata {

		documents[i] = userdata
	}

	_, err := m.UsersCollection.InsertMany(context.TODO(), documents)
	if err != nil {
		return fmt.Errorf("failed to store Data in the database: %v", err)
	}

	return nil
}

func (m *MongoStore) QueryRecentActions() ([]models.RecentActions, error) {
	var actions []models.RecentActions

	cursor, err := m.RecentActionsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to query recentActions from the database: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.Background()) {
		var action models.RecentActions
		if err := cursor.Decode(&action); err != nil {
			return nil, fmt.Errorf("failed to decode recentActions document: %v", err)
		}
		actions = append(actions, action)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while querying recentActions from the database: %v", err)
	}

	return actions, nil
}
func (m *MongoStore) QueryUsers() ([]models.Users, error) {
	var userdata []models.Users

	cursor, err := m.UsersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to query users from the database: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.Background()) {
		var userdat models.Users
		if err := cursor.Decode(&userdat); err != nil {
			return nil, fmt.Errorf("failed to decode Users document: %v", err)
		}
		userdata = append(userdata, userdat)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while querying Users from the database: %v", err)
	}

	return userdata, nil
}

// func (m *MongoStore) QueryUsersByEmail(email string) ([]models.Users, error) {

//		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
//		var users []models.Users
//		cursor, err := m.UsersCollection.Find(ctx, bson.M{"email": email})
//		if err != nil {
//			return nil, fmt.Errorf("failed to query users from the database: %v", err)
//		}
//		if err := cursor.Err(); err != nil {
//			return nil, fmt.Errorf("cursor error while querying Users from the database: %v", err)
//		}
//		defer cancel()
//		fmt.Println(users)
//		return users, nil
//	}
func (m *MongoStore) Close() {
	if err := m.RecentActionsCollection.Database().Client().Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	}
}
func OpenCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(databaseName).Collection(collectionName)

	return collection
}
