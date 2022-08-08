package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Travel struct {
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

func GetItemsTransportedByUser(Id int64) []Travel {
	var itemsTransported []Travel
	db.Where("UserId=?", Id).Find(&itemsTransported)
	return itemsTransported
}
