package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"example.com/go-server/database"
	"example.com/go-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProduct handles saving a product to MongoDB
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product // getting the product model

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set unique ID
	product.ID = primitive.NewObjectID()

	// Insert into MongoDB
	collection := database.Client.Database("mydb").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, product)
	if err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product created successfully"})
}

// GetProducts retrieves all products from the database
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Set a timeout for database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// finding the collection
	collection := database.Client.Database("mydb").Collection("products")

	// Find all documents (empty filter `{}` fetches all)
	cursor, err := collection.Find(ctx, bson.M{}) // cursor holds the data
	if err != nil {
		http.Error(w, "Failed to get products from database", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx) // Close cursor after function returns

	// Slice to store retrieved products
	var products []models.Product

	// Iterate over the cursor and decode each document
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			http.Error(w, "Error decoding product data", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	// Check if no products were found
	if len(products) == 0 {
		http.Error(w, "No products found", http.StatusNotFound)
		return
	}

	// Set response headers and encode JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
