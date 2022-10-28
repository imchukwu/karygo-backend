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

// Enable access from all origin
func enableCors(w *http.ResponseWriter) {
(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
// swagger:route GET /admin admins listAdmins
// 	Returns a list of admins
// 	Responses:
//  	200: adminsResponse
//  Produces:
// 		- application/json
//		- application/xml

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

// swagger:route POST /admin admin creatAdmin
// 	Creates new admin
// 	responses:
//  	200: createAdminResponse
// 		
// 		default: error
//  Produces:
// 		- application/json
//		- application/xml

// New admin to be returned in response
// swagger:response createAdminResponse
type createAdminResponse struct{
	// Create new admin in the system
	// in: body
	Body models.Admin

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

// swagger:route PUT /admin/{adminId} admin updateAdmin 
// 	Returns updated admin
// 	Parameters:
// 		name: adminId
// 		in: path
// 		description: ID of admin to be updated
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: updateAdminResponse
//  Produces:
// 		- application/json
//		- application/xml

// updated admin is returned in response
// swagger:response updateAdminResponse
type updateAdminResponse struct{
	// update admin in the system
	// in: body
	Body models.Admin

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

// swagger:route DELETE /admin/{adminId} admin deleteAdmin
// 	Returns deleted admin
// 	Parameters:
// 		name: adminId
// 		in: path
// 		description: ID of admin to be deleted
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: deleteAdminResponse
//  Produces:
// 		- application/json
//		- application/xml

// deleted admin is returned in response
// swagger:response deleteAdminResponse
type deleteAdminResponse struct{
	// delete admin in the system
	// in: body
	Body models.Admin

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
// swagger:route GET /admin/{adminId} admin getAdmin
// 	Returns an admin by ID
// 	Parameters:
// 		name: adminId
// 		in: path
// 		description: ID of admin to be fetched
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: adminResponse
//  Produces:
// 		- application/json
//		- application/xml

// admin to be returned in response
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

// swagger:route GET /admin/transactions transaction getTransactions
// Returns list of transactions
// 	Responses:
//  	200: transactionsResponse
//  Produces:
// 		- application/json
//		- application/xml


// list of transactions to be returned in response
// swagger:response transactionsResponse
type transactionsResponse struct{
	// All Transactions in the system
	// in: body
	Body []models.Transaction

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

// swagger:route GET /admin/transaction/{txId} transaction getTransaction
// Returns a transaction by ID
// 	Parameters:
// 		in: path
// 		required: true 
// 		name: txId
// 		description: ID of transaction to be fetched
// 		
//		type: integer
// 		format: int64
// 		
// 	Responses:
//  	200: transactionResponse
//  Produces:
// 		- application/json
//		- application/xml

// transaction to be returned in response
// swagger:response transactionResponse
type transactionResponse struct{
	// A transaction in the system by ID
	// in: body
	Body models.Transaction

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