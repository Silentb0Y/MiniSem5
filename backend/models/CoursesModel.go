package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Image struct {
	Data []byte `bson:"data" json:"data"`
}

type Course struct {
	CourseID            int      `bson:"_id,omitempty" json:"courseId"`
	CourseName          string   `bson:"courseName" json:"courseName"`
	CourseMetaData      string   `bson:"courseMetaData" json:"courseMetaData"`
	CourseResourceLinks []string `bson:"courseResourceLinks" json:"courseResourceLinks"`
	CourseImage         Image    `bson:"courseImage" json:"courseImage"`
}

var courseCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("codeforprogress")
	courseCollection = database.Collection("courses")
}

func InsertOneCourse(course Course) {
	inserted, err := courseCollection.InsertOne(context.Background(), course)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted one course with ID:", inserted.InsertedID)
}

func GetCoursesMetaData() []Course {
	var courses []Course

	cursor, err := courseCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var course Course
		if err := cursor.Decode(&course); err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}
	return courses
}

func GetCourseData(courseId int) Course {
	var course Course
	err := courseCollection.FindOne(context.TODO(), bson.M{"courseId": courseId}).Decode(&course)
	if err == mongo.ErrNoDocuments {
		fmt.Println("Course not found")
	} else if err != nil {
		log.Fatal(err)
	}
	return course
}
