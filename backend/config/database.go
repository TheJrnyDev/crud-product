package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// global variables to hold the MongoDB client and database
var (
	Client   *mongo.Client
	Database *mongo.Database
)

func InitDatabase() *mongo.Client {
	// create a context with a timeout for the database connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// cancel the context when done to free resources
	defer cancel()

	// get connection string from environment variable, fallback to default
	connectionString := os.Getenv("MONGODB_URI")
	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
	}

	// get database name from environment variable, fallback to default
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		databaseName = "crud-product"
	}

	// connect to the database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println("❌ Failed to connect to MongoDB:")
		log.Fatal(err)
	}

	// ping the database to check if the connection is successful
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("❌ Failed to ping MongoDB:")
		log.Fatal(err)
	}

	// Store globally for use throughout the app
	Client = client
	Database = client.Database(databaseName)

	fmt.Printf("✅ Connected to MongoDB database: %s\n", databaseName)

	return Client
}

func CloseDatabase() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		Client.Disconnect(ctx)
		fmt.Println("✅ Database connection closed")
	}
}
