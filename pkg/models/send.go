package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Send struct {
	gorm.Model
	UserID uint  `json:"userId"`
	ItemID uint  `json:"itemId"`
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
