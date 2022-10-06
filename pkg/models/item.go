package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id              uint   `json:"id"`
	UserId          uint   `json:"userId"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	TransctionId    uint   `json:"transactionId"`
	SuccessCode     string `json:"success_code"`
	TripID          uint   `json:"tripId"`
	Trip            *Trip  `json:"trip" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Item{})
}

func (item *Item) CreateItem() *Item {
	db.Create(&item)
	return item
}

func GetItems() []Item {
	var items []Item
	db.Find(&items)
	return items
}

func GetItem(Id int64) (*Item, *gorm.DB) {
	var item Item
	db := db.Where("ID=?", Id).Find(&item)
	return &item, db
}

func GetItemsByTrip(Id int64) []Item {
	var items []Item
	db.Where("ID=?", Id).Find(&items)
	return items
}

func DeleteItem(Id int64) *Item {
	var item *Item
	db.Where("ID=?", Id).Delete(&item)
	return item
}

// func DeleteItemsByUser(Id int64) []*Item {
// 	var items []*Item
// 	db.Where("ID=?", Id).Delete(&items)
// 	return items
// }

func DeleteItemsByTrip(Id int64) []*Item {
	var items []*Item
	db.Where("ID=?", Id).Delete(&items)
	return items
}
