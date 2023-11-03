package main

import (
	"coursemicro/som"
	"log"
)

func main() {
	store, err := som.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	server := som.NewAPIServer(":7002", store)
	server.Run()
}
