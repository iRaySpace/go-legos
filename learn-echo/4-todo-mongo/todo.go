package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name"`
}
