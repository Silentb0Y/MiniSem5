package controller

import (
	model "backboo/models"
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var data model.ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// You can send a response back to the client if needed
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("JSON data received and processed successfully \n"))
}

func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData model.ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}

	courseslist := model.GetCoursesMetaData()
	w.WriteHeader(http.StatusOK)
	outgoingData := model.ApiData{
		Courselist: courseslist,
	}
	marsheledOutgoingData, err := json.Marshal(outgoingData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(marsheledOutgoingData)
}

func GetCourseHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData model.ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}

	course := model.GetCourseData(incomingData.CourseId)
	w.WriteHeader(http.StatusOK)
	outgoingData := model.ApiData{
		Course: course,
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
