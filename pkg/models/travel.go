package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model
	ID uint  `json:"id"`
	UserID uint  `json:"userId"`
	ItemID uint  `json:"itemId"`
	TravelId uint  `json:"travelId"`
	TripId uint  `json:"tripId"`
	SuccessCode string  `json:"successCode"`
	//followed by SendDate field whose values to be generated using timestamp
	Item   *Item `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Item{})
}

func GetItemsTransportedByUser(Id int64) []Travel {
	var itemsTransported []Travel
	db.Where("UserId=?", Id).Find(&itemsTransported)
	return itemsTransported
}
