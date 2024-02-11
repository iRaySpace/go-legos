package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongo() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Unable to connect, MONGO_URI is not set.")
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

func DisconnectMongo() {
	if client == nil {
		return
	}
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	client = nil
}
