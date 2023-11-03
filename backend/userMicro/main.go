package main

import (
	"log"
	"usermicro/som"
)

func main() {
	store, err := som.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	server := som.NewAPIServer(":7001", store)
	server.Run()
}

// func main() {
// 	router := mux.NewRouter()
// 	corsHandler := gohandlers.CORS(
// 		gohandlers.AllowedOrigins([]string{"*"}),
// 		gohandlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE"}),                  // Allow only GET and POST methods
// 		gohandlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}), // Allow only Content-Type header
// 	)

// 	router.HandleFunc("/login", doshit)
// 	router.HandleFunc("/signup", doshit)

// 	log.Println("JSON API server running on port:", ":9000")
// 	http.ListenAndServe(":7001", corsHandler(router))
// }
// func doshit(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("shit")
// }

// func d(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("workn")
// }
