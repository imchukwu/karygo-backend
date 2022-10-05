package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Id              uint   `json:"id"`
	UserId          uint   `json:"userId"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	ReceiverPhone   string `json:"receiver_phone"`
	TransctionType  string `json:"transactionId"`
	SuccessCode     string `json:"success_code"`
	TripID          uint   `json:"tripId"`
	Trip            *Trip  `json:"trip" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

// func GetItemsByTrip(Id int64) []Item {
// 	var items []Item
// 	db.Where("ID=?", Id).Find(&items)
// 	return items
// }

// func DeleteItem(Id int64) *Item {
// 	var item *Item
// 	db.Where("ID=?", Id).Delete(&item)
// 	return item
// }

// // func DeleteItemsByUser(Id int64) []*Item {
// // 	var items []*Item
// // 	db.Where("ID=?", Id).Delete(&items)
// // 	return items
// // }

// func DeleteItemsByTrip(Id int64) []*Item {
// 	var items []*Item
// 	db.Where("ID=?", Id).Delete(&items)
// 	return items
// }
