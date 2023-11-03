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
	router.HandleFunc("/getCourses", s.GetCoursesHandler)
	router.HandleFunc("/getcourse/:id", s.GetCourseHandler)

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, corsHandler(router))
}

func (s *APIServer) GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	courseslist, _ := s.store.GetAllCourses()
	w.WriteHeader(http.StatusOK)
	outgoingData := ApiData{
		Courselist: *courseslist,
	}
	marsheledOutgoingData, err := json.Marshal(outgoingData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("sending courses")
	w.Write(marsheledOutgoingData)
}

func (s *APIServer) GetCourseHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}

	course, _ := s.store.GetCourseData(incomingData.CourseId)
	w.WriteHeader(http.StatusOK)
	outgoingData := ApiData{
		Course: *course,
	}
	marsheledOutgoingData, err := json.Marshal(outgoingData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(marsheledOutgoingData)
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Page Not Found"))
}
