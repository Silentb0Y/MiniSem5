package main

import (
	controller "backboo/controllers"

	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/codeforprogress"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create a new router
	router := mux.NewRouter()

	// Define routes and their corresponding handlers
	router.HandleFunc("/", controller.HomeHandler)
	router.HandleFunc("/login", controller.LoginHandler)
	router.HandleFunc("/sigup", controller.SignupHandler)
	router.HandleFunc("/getCourses", controller.GetCoursesHandler)
	router.HandleFunc("/getCourse", controller.GetCourseHandler)

	// Custom 404 handler
	router.NotFoundHandler = http.HandlerFunc(controller.NotFoundHandler)

	// Start the HTTP server
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
