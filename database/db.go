package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Initialize MongoDB Connection
func ConnectDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load environment variables")
	}

	// Get MongoDB URL
	mongoURL := os.Getenv("MONGODB_URI")
	if mongoURL == "" {
		log.Fatal("MongoDB URI is missing")
	}
	// log.Fatalf(mongoURL)
	// Set connection options
	clientOptions := options.Client().ApplyURI(mongoURL)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errDB error // to connect to the server
	Client, errDB = mongo.Connect(ctx, clientOptions)
	if errDB != nil {
		log.Fatal("MongoDB connection error:", errDB)
	}

	// Ping the database to check the connection
	errPing := Client.Ping(ctx, nil)
	if errPing != nil {
		log.Fatal("MongoDB ping error:", errPing)
	}

	fmt.Println("✅ Connected to MongoDB successfully!")
}
