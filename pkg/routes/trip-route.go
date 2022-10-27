package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/controllers"
)

var RegisterTripRoutes = func(r *mux.Router) {
	r.HandleFunc("/trips", controllers.CreateTrip).Methods("POST")
	r.HandleFunc("/trips", controllers.GetTrips).Methods("GET")
	r.HandleFunc("/trip/{tripId}", controllers.GetTrip).Methods("GET")
	r.HandleFunc("/trips/{userId}/complete", controllers.GetCompletedTripsByUser).Methods("GET")
	r.HandleFunc("/trips/{userId}/pending", controllers.GetPendingTripsByUser).Methods("GET")
	// r.HandleFunc("/trips/{userId}", controllers.GetTripsByUser).Methods("GET")
	// r.HandleFunc("/trip/{userId}/{tripId}", controllers.UpdateTrip).Methods("PUT")
	r.HandleFunc("/trip/{tripId}", controllers.DeleteTrip).Methods("DELETE")
	// r.HandleFunc("/trip/{userId}", controllers.DeleteTripsByUser).Methods("DELETE")
}
