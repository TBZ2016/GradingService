package dal

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

// Initialize initializes the MongoDB client and connects to the database.
func Initialize(connectionString, dbName string) error {
	clientOptions := options.Client().ApplyURI(connectionString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	database = client.Database(dbName)
	return nil
}

// GetDatabase returns the MongoDB database instance.
func GetDatabase() *mongo.Database {
	return database
}

// Close closes the MongoDB client when it's no longer needed.
func Close() {
	if client != nil {
		client.Disconnect(context.Background())
	}
}
