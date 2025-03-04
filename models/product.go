package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product struct to store in MongoDB
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Price       float64            `bson:"price,omitempty"`
	Description string             `bson:"description,omitempty"`
}
