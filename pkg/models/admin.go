package models

import (
	"github.com/imchukwu/karygo_backend/pkg/config"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	AdminType  string `json:"admin_type"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Admin{})
}

func (adminUser *Admin) CreateAdmin() *Admin {
	db.Create(&adminUser)
	return adminUser
}

func GetAdmins() []Admin {
	var admins []Admin
	db.Find(&admins)
	return admins
}

func GetAdmin(Id int64) (*Admin, *gorm.DB) {
	var admin Admin
	db := db.Where("ID=?", Id).Find(&admin)
	return &admin, db
}

func DeleteAdmin(Id int64) *Admin {
	var admin *Admin
	db.Where("ID=?", Id).Delete(&admin)
	return admin
}
