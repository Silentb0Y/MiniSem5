package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string
	Password string
}

var userCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("codeforprogress")
	userCollection = database.Collection("courses")
}

func GetUsernameAndPassword(username string) User {
	var result User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("Course not found")
	} else if err != nil {
		log.Fatal(err)
	}
	return result
}

func AddUser(user User) {
	inserted, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted one user with ID:", inserted.InsertedID)
}
