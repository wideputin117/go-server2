package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"example.com/go-server/database"
	"example.com/go-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProduct handles saving a product to MongoDB
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

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
