package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	DepartureState        string `json:"departure_state"`
	DepartureLocation     string `json:"departure_location"`
	DepartureDate         string `json:"departure_date"`
	DepartureTime         string `json:"departure_time"`
	ArrivalState          string `json:"arrival_state"`
	ArrivalLocation       string `json:"arrival_location"`
	MeansOfTransportation string `json:"means_of_transportation"`
	TravelerId            uint   `json:"travelerId"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Trip{})
}

func (trip *Trip) CreateTrip() *Trip {
	db.Create(&trip)
	return trip
}

func GetTrips() []Trip {
	var trips []Trip
	db.Find(&trips)
	return trips
}

func GetTrip(Id int64) (*Trip, *gorm.DB) {
	var trip Trip
	db := db.Where("ID=?", Id).Find(&trip)
	return &trip, db
}

func DeleteTrip(Id int64) *Trip {
	var trip *Trip
	db.Where("ID=?", Id).Delete(&trip)
	return trip
}
