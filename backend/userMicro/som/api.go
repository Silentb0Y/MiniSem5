package som

import (
	"encoding/json"
	"log"
	"net/http"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	corsHandler := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE"}),                  // Allow only GET and POST methods
		gohandlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}), // Allow only Content-Type header
	)
	router.HandleFunc("/login", s.LoginHandler)
	router.HandleFunc("/signup", s.SignupHandler)
	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, corsHandler(router))
}

func (s *APIServer) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}
	s.store.CreateUserInDB(incomingData.Username, incomingData.Password)
	w.WriteHeader(http.StatusOK)
}

func (s *APIServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}

	tureUser, err := s.store.CheckUservalidityAndGetUser(incomingData.Username, incomingData.Password)

	if !tureUser {
		log.Println("User sent worng credentials")
		http.Error(w, "Wrong credentials", http.StatusBadRequest)
		return
	}

	log.Println("user loggedIn with credentials", incomingData.Username, incomingData.Password)

	w.Header().Set("Authorization", "Bearer "+"token")
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("login sucessful"))
}
