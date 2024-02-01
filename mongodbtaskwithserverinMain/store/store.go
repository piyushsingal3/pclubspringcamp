package store

import (
	"context"
	"fmt"
	"log"

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
	fmt.Print("Connection Establ")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	db := client.Database(databaseName)

	m.RecentActionsCollection = db.Collection("recentActions")
	m.UsersCollection = db.Collection("users")

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

func (m *MongoStore) QueryRecentActions() ([]models.RecentActions, error) {
	var actions []models.RecentActions

	cursor, err := m.RecentActionsCollection.Find(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query recentActions from the database: %v", err)
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &actions)
	if err != nil {
		return nil, fmt.Errorf("failed to decode recentActions documents: %v", err)
	}

	return actions, nil
}

func (m *MongoStore) Close() {
	if err := m.RecentActionsCollection.Database().Client().Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	}
}