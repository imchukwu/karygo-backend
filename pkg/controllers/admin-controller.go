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

// Enable access from localhost:3000
func enableCors(w *http.ResponseWriter) {
(*w).Header().Set("Access-Control-Allow-Origin", "*")
(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS") 
(*w).Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")

}
// swagger:route GET /admin admins listAdmins
// 	Returns a list of admins
// 	Responses:
//  	200: adminsResponse


// list of admins to be returned in response
// swagger:response adminsResponse
type adminsResponse struct{
	// All admins in the system
	// in: body
	Body []models.Admin

}

// Get all Admins
func GetAdmins (w http.ResponseWriter, r *http.Request){
	enableCors(&w)

	admins := models.GetAdmins()
	res, _ := json.Marshal(admins)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Create a new Admin
func CreateAdmin (w http.ResponseWriter, r *http.Request){
	enableCors(&w)

	newAdmin := &models.Admin{}
	utils.ParseBody(r, newAdmin)
	admin := newAdmin.CreateAdmin()
	res, _ := json.Marshal(admin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Update an existing Admin by parameter:ID
func UpdateAdmin (w http.ResponseWriter, r *http.Request){
	enableCors(&w)

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

// Delete Admin by parameter:ID
func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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
// swagger:route GET /admin/{adminId} admin admin
// 	Returns an admin by ID
// 	Parameters:
// 		name: // 	responses:
//  	200: adminResponse

// list of admins to be returned in response
// swagger:response adminResponse
type adminResponse struct{
	// All admins in the system
	// in: body
	Body models.Admin

}

// Get admin by parameter:ID
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	adminId := vars["adminId"]
	Id, err := strconv.ParseInt(adminId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	adminDetail, _ := models.GetAdmin(Id)
	res, _ := json.Marshal(adminDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route GET /admin/{adminId} admin admin
// Returns a list of admins
// Responses:
//  200: adminResponse

// list of admins to be returned in response
// swagger:response adminResponse
type transactionsResponse struct{
	// All admins in the system
	// in: body
	Body models.Admin

}

// Get all Transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	transactions := models.GetTransactions()
	res, _ := json.Marshal(transactions)
	w.Header().Set("Content_Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Get transaction by parameter:ID
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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