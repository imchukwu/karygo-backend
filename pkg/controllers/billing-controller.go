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

func GetBillings(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	billings := models.GetBillings()
	res, _ := json.Marshal(billings)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBilling(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	billingId := vars["billingId"]
	Id, err := strconv.ParseInt(billingId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	billingDetail, _ := models.GetBilling(Id)
	res, _ := json.Marshal(billingDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateBilling(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	createBilling := &models.Billing{}
	utils.ParseBody(r, createBilling)
	u := createBilling.CreateBilling()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBilling(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var UpdateBilling = &models.Billing{}
	utils.ParseBody(r, UpdateBilling)

	vars := mux.Vars(r)
	billingId := vars["billingId"]
	Id, err := strconv.ParseInt(billingId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	billingDetail, db := models.GetBilling(Id)
	if UpdateBilling.UserId != 0 {
		billingDetail.UserId = UpdateBilling.UserId
	}
	if UpdateBilling.PreviousBalance != 0 {
		billingDetail.PreviousBalance = UpdateBilling.PreviousBalance
	}
	if UpdateBilling.CurrentBalance != 0 {
		billingDetail.CurrentBalance = UpdateBilling.CurrentBalance
	}
	if UpdateBilling.PaymentType != "" {
		billingDetail.PaymentType = UpdateBilling.PaymentType
	}
	if UpdateBilling.Amount != 0 {
		billingDetail.Amount = UpdateBilling.Amount
	}
	if UpdateBilling.Status != "" {
		billingDetail.Status = UpdateBilling.Status
	}


	db.Save(&billingDetail)

	res, _ := json.Marshal(billingDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBilling(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	billingId := vars["billingId"]
	Id, err := strconv.ParseInt(billingId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	billing := models.DeleteBilling(Id)
	res, _ := json.Marshal(billing)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBillingsByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	var billings []*models.Billing
	billings = models.GetBillingsByUser(Id)
	res, _ := json.Marshal(billings)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}