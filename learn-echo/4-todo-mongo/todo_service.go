package main

import (
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct{}

func (ts *TodoService) Create(t *Todo) *Todo {
	collection := client.Database(os.Getenv("MONGO_DB")).Collection("todos")

	newT := Todo{
		ID:   primitive.NewObjectID(),
		Name: t.Name,
	}

	_, err := collection.InsertOne(context.TODO(), newT)
	if err != nil {
		log.Fatal(err)
	}

	return &newT
}

func (ts *TodoService) GetAll() []*Todo {
	collection := client.Database(os.Getenv("MONGO_DB")).Collection("todos")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var todos []*Todo
	for cur.Next(context.TODO()) {
		var t Todo
		err := cur.Decode(&t)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, &t)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return todos
}

func (ts *TodoService) Update(id string, t *Todo) error {
	collection := client.Database(os.Getenv("MONGO_DB")).Collection("todos")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	newT := Todo{
		ID:   idPrimitive,
		Name: t.Name,
	}

	res, err := collection.UpdateOne(context.TODO(), bson.M{"_id": newT.ID}, bson.M{"$set": newT})
	if err != nil {
		log.Fatal(err)
	}

	if res.MatchedCount == 0 {
		return errors.New("no document has found")
	}

	return nil
}

func (ts *TodoService) Delete(id string) error {
	collection := client.Database(os.Getenv("MONGO_DB")).Collection("todos")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		log.Fatal(err)
	}

	if res.DeletedCount == 0 {
		return errors.New("no document has found")
	}

	return nil
}
