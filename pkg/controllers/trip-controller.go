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

// swagger:route GET /trips trips listTrips
// 	Returns a list of trips
// 	Responses:
//  	200: tripsResponse
//  Produces:
// 		- application/json
//		- application/xml

// list of trips to be returned in response
// swagger:response tripsResponse
type tripsResponse struct{
	// All trips in the system
	// in: body
	Body []models.Trip

}


// Get all trips
func GetTrips(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	trips := models.GetTrips()
	res, _ := json.Marshal(trips)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route GET /trip/{tripId} trip gettrip
// 	Returns an trip by ID
// 	Parameters:
// 		name: tripId
// 		in: path
// 		description: ID of trip to be fetched
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: tripResponse
//  Produces:
// 		- application/json
//		- application/xml

// trip to be returned in response
// swagger:response tripResponse
type tripResponse struct{
	// All trips in the system
	// in: body
	Body models.Trip

}


// Get trip by ID
func GetTrip(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

// swagger:route POST /trips trip creatTrip
// 	Creates new trip
// 	responses:
//  	200: createTripResponse
// 		
// 		default: error
//  Produces:
// 		- application/json
//		- application/xml

// New trip to be returned in response
// swagger:response createtripResponse
type createtripResponse struct{
	// Create new trip in the system
	// in: body
	Body models.Trip

}


// Creates new trip 
func CreateTrip(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	CreateTrip := &models.Trip{}
	utils.ParseBody(r, CreateTrip)
	u := CreateTrip.CreateTrip()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route GET /trips/{userId}/complete trip getCompletedTripsByUser
// Returns list of completed trips by userId
// 	Parameters:
// 		in: path
// 		required: true 
// 		name: userId
// 		description: ID of user to be used in fetching completed trips
// 		
//		type: integer
// 		format: int64
// 		
// 	Responses:
//  	200: completedTripsResponse
//  Produces:
// 		- application/json
//		- application/xml

// Trips to be returned in response
// swagger:response completedTripsResponse
type completedTripsResponse struct{
	// A transaction in the system by ID
	// in: body
	Body []models.Trip

}


// Gets completed trips by user
func GetCompletedTripsByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

// swagger:route GET /trips/{userId}/pending trip getPendingTripsByUser
// Returns list of pending trips by userId
// 	Parameters:
// 		in: path
// 		required: true 
// 		name: userId
// 		description: ID of user to be used in fetching pending trips
// 		
//		type: integer
// 		format: int64
// 		
// 	Responses:
//  	200: pendingTripsResponse
//  Produces:
// 		- application/json
//		- application/xml

// Trips to be returned in response
// swagger:response pendingTripsResponse
type pendingTripsResponse struct{
	// A transaction in the system by ID
	// in: body
	Body []models.Trip

}

// Gets pending trips by user
func GetPendingTripsByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

// swagger:route DELETE /trip/{tripId} trip deleteTrip
// 	Returns deleted trip
// 	Parameters:
// 		name: tripId
// 		in: path
// 		description: ID of trip to be deleted
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: deleteTripResponse
//  Produces:
// 		- application/json
//		- application/xml

// deleted trip is returned in response
// swagger:response deleteTripResponse
type deleteTripResponse struct{
	// delete trip in the system
	// in: body
	Body models.Trip

}

// Deletes trip from system by Id
func DeleteTrip(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	
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
