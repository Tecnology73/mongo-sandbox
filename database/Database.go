package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db            *mongo.Database
	isInitialized bool
}

var dbInstance *Database

func Collection(name string) *mongo.Collection {
	return Instance().db.Collection(name)
}

func Instance() *Database {
	if dbInstance == nil {
		// Initiate the database
		fmt.Printf("[Database] Initializing...\n")

		dbInstance = new(Database)

		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://user:mongoatlas@sandbox.ah5kk.mongodb.net/sandbox?w=majority").
			SetServerAPIOptions(serverAPIOptions).
			SetRetryWrites(true)

		// Make connect return an error if it doesn't complete within 10 seconds
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		// Connect
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Verify the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		// Select the database
		dbInstance.db = client.Database("sandbox")
		dbInstance.isInitialized = true
	}

	return dbInstance
}
