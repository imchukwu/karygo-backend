package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/models"
	"github.com/imchukwu/karygo_backend/pkg/utils"
)

func GetTrips(w http.ResponseWriter, r *http.Request) {
	trips := models.GetTrips()
	res, _ := json.Marshal(trips)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripId := vars["tripId"]
	Id, err := strconv.ParseInt(tripId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	tripDetail, _ := models.GetTrip(Id)
	res, _ := json.Marshal(tripDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTrip(w http.ResponseWriter, r *http.Request) {
	CreateTrip := &models.Trip{}
	utils.ParseBody(r, CreateTrip)
	u := CreateTrip.CreateTrip()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCompletedTripsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	var tripIds []uint
	var trips []*models.Trip
	tripIds = models.GetCompletedTripByUser(Id)
	for index := range tripIds {
		trip, _ := models.GetTrip(Id)
		trips[index] = trip
	}

	res, _ := json.Marshal(trips)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPendingTripsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	var tripIds []uint
	var trips []*models.Trip
	tripIds = models.GetPendingTripByUser(Id)
	for index := range tripIds {
		trip, _ := models.GetTrip(Id)
		trips[index] = trip
	}

	res, _ := json.Marshal(trips)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripId := vars["tripId"]
	Id, err := strconv.ParseInt(tripId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	trip := models.DeleteTrip(Id)
	res, _ := json.Marshal(trip)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
