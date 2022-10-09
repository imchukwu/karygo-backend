package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TravelerId  uint   `json:"travelerId"`
	SenderId    uint   `json:"senderId"`
	ItemId      uint   `json:"itemId"`
	TripId      uint   `json:"tripId"`
	SuccessCode string `json:"success_code"`
	Status      string `json:"status"` //Complete, Pending, Failed, Ongoing
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Transaction{})
}

func (transaction *Transaction) CreateTransaction() *Transaction {
	db.Create(&transaction)
	return transaction
}

func GetTransactions() []Transaction {
	var transactions []Transaction
	db.Find(&transactions)
	return transactions
}

func GetTransaction(Id int64) (*Transaction, *gorm.DB) {
	var transaction Transaction
	db := db.Where("ID=?", Id).Find(&transaction)
	return &transaction, db
}

func GetItemIdBySender(Id int64) []uint {
	var itemIds []uint
	db.Where("SenderId=?", Id).Find(itemIds)
	return itemIds
}

func GetItemIdByTraveler(Id int64) []uint {
	var itemIds []uint
	db.Where("TravelerId=?", Id).Find(itemIds)
	return itemIds

}

func GetCompletedTripByUser(Id int64) []uint {
	var tripIds []uint
	db.Where("Status = ?", "Complete").Where("SenderId = ?", Id).Or("TravelerId = ?", Id).Find(tripIds)
	return tripIds
}

func GetPendingTripByUser(Id int64) []uint {
	var tripIds []uint
	db.Where("Status = ?", "Pending").Where("SenderId = ?", Id).Or("TravelerId = ?", Id).Find(tripIds)
	return tripIds

}

func DeleteTransaction(Id int64) *Transaction {
	var transaction *Transaction
	db.Where("ID=?", Id).Delete(&transaction)
	return transaction
}
