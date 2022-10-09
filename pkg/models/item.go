package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	SenderId                  uint   `json:"senderId"`
	Name                      string `json:"name"`
	ItemImageId               uint   `json:"item_imageId"`
	Start                     string `json:"start"`
	End                       string `json:"end"`
	DateRange                 string `json:"date_range"`
	PreferredMeansOfTransport string `json:"preferred_means_of_transport"`
	ReceiverName              string `json:"receiver_name"`
	ReceiverPhone             string `json:"receiver_phone"`
	ReceiverImageId           uint   `json:"receiver_imageId"`
	Quantity                  uint   `json:"quantity"`
	ItemImage                 *Image `json:"item_image" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReceiverImage             *Image `json:"receiver_image" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

func DeleteItem(Id int64) *Item {
	var item *Item
	db.Where("ID=?", Id).Delete(&item)
	return item
}
