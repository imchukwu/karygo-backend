package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Send struct {
	gorm.Model
	Id     uint  `json:"id"`
	UserID uint  `json:"userId"`
	ItemID uint  `json:"itemId"`
	ReceiverPhone   string `json:"receiver_phone"`
	ReceivedDate  string `json:"received_date_time"`
	SuccessCode string  `json:"success_code"`
	//followed by SendDate field whose values to be generated using timestamp
	Item   *Item `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Item{})
}

func GetItemsSentByUser(Id int64) []Send {
	var itemSent []Send
	db.Where("UserId=?", Id).Find(&itemSent)
	return itemSent
}

func DeleteItemsByUser(Id int64) []Send {
	var sends []Send
	db.Where("UserId=?", Id).Find(&sends)

	for _, item := range sends {
		db.Delete(item.Item)
	}
	return sends
}
