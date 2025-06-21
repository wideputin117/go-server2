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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// lets marsahl the oncoming body
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Inavlid request payload", http.StatusBadRequest)
		return
	}
	// set unique id
	user.ID = primitive.NewObjectID()
	// lets get the instance of the database for inserting the data
	collection := database.Client.Database("mydb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
