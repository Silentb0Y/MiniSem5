package controller

import (
	model "backboo/models"
	"encoding/json"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData model.ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}
	model.AddUser(model.User{
		Username: incomingData.Username,
		Password: incomingData.Password,
	})
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData model.ApiData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}

	userdata := model.GetUsernameAndPassword(incomingData.Username)

	if userdata.Username != incomingData.Username && userdata.Password != incomingData.Password {
		http.Error(w, "Invalid", http.StatusBadRequest)
		return
	}

	w.Header().Set("Authorization", "Bearer "+"token")
	w.WriteHeader(http.StatusOK)

}
