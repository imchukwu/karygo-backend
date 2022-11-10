// Package classification of admin API
//
// Documentation for admin API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/models"
	"github.com/imchukwu/karygo_backend/pkg/utils"
)

func GetAdmins (w http.ResponseWriter, r *http.Request){
	admins := models.GetAdmins()
	res, _ := json.Marshal(admins)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAdmin (w http.ResponseWriter, r *http.Request){
	newAdmin := &models.Admin{}
	utils.ParseBody(r, newAdmin)
	admin := newAdmin.CreateAdmin()
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAdmin (w http.ResponseWriter, r *http.Request){
	updateAdmin := &models.Admin{}
	utils.ParseBody(r, updateAdmin)

	vars := mux.Vars(r)
	adminId := vars["adminId"]
	Id, err := strconv.ParseInt(adminId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	adminDetail, db := models.GetAdmin(Id)
	if updateAdmin.FirstName != "" {
		adminDetail.FirstName = updateAdmin.FirstName
	}
	if updateAdmin.LastName != "" {
		adminDetail.LastName = updateAdmin.LastName
	}
	if updateAdmin.Email != "" {
		adminDetail.Email = updateAdmin.Email
	}
	if updateAdmin.Phone != "" {
		adminDetail.Phone = updateAdmin.Phone
	}
	if updateAdmin.AdminType != "" {
		adminDetail.AdminType = updateAdmin.AdminType
	}
	if updateAdmin.Department != "" {
		adminDetail.Department = updateAdmin.Department
	}
	if updateAdmin.JobTitle != "" {
		adminDetail.JobTitle = updateAdmin.JobTitle
	}

	db.Save(&adminDetail)

	res, _ := json.Marshal(adminDetail)
	w.Header().Set("Content-Type", "applicatiion/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminId := vars["adminId"]
	Id, err := strconv.ParseInt(adminId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	admin := models.DeleteAdmin(Id)
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminId := vars["adminId"]
	Id, err := strconv.ParseInt(adminId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing me")
	}
	adminDetail, _ := models.GetAdmin(Id)
	res, _ := json.Marshal(adminDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := models.GetTransactions()
	res, _ := json.Marshal(transactions)
	w.Header().Set("Content_Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txId := vars["txId"]
	Id, err := strconv.ParseInt(txId, 0, 0)
	if err != nil {
		fmt.Println("error while passing")
	}

	txDetail, _ := models.GetTransaction(Id)
	res, _ := json.Marshal(txDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}