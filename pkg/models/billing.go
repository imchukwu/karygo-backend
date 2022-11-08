package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Billing struct {
	gorm.Model

	UserId  		  uint  `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PreviousBalance   uint   `json:"previousBalance"`
	CurrentBalance	  uint   `json:"currentBalance"`
	PaymentType       string `json:"paymentType"`
	Amount 			  uint	 `json:"amount"`
	Status			  string `json:"status"`
}


	func init(){
		config.Connect()
		db = config.GetDB()
		db.AutoMigrate(&Billing{})
	}


	func (billing *Billing) CreateBilling() *Billing{
		db.Create(&billing)
		return billing
	}


	func GetBillings() []Billing{
		var billings []Billing
		db.Find(&billings)
		return billings
	}


	func  GetBilling(Id int64) (*Billing, *gorm.DB){
		var billing Billing
		db := db.Where("ID=?", Id).Find(&billing)
		return &billing, db
	}

	func GetBillingsByUser(Id int64) []uint{
		var userBillings []uint
		db.Where("userId = ?", Id).Find(userBillings)
		return userBillings
	}

	// func GetDebitBillingsByUser(Id int64) []Billings{
	// 	 var userBillings []Billings
	// 	 db.Where("UserId = ?", Id).Where("TransactionType = ?", "Debit").Find(&userBillings)
	// 	 return userBillings
	// }

	// func GetCreditBillingsByUser(Id uint64) []Billings{
	// 	var userBillings []Billings
	// 	db.Where("UserId = ?", Id).Where("TransactionType = ?", "Credit").Find(&userBillings)
	// 	return userBillings
	// }

	func DeleteBilling(Id int64) *Billing{
		var billing *Billing
		db.Where("ID = ?", Id).Delete(&billing)
		return billing
	}
