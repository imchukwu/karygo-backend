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

func GetItems(w http.ResponseWriter, r *http.Request) {
	items := models.GetItems()
	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	Id, err := strconv.ParseInt(itemId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	itemDetail, _ := models.GetItem(Id)
	res, _ := json.Marshal(itemDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetItemsSentByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	var items []*models.Item
	sends := models.GetItemsSentByUser(Id)
	for index, item := range sends {
		items[index] = item.Item
	}
	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetItemsTransportedByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	var items []*models.Item
	travels := models.GetItemsTransportedByUser(Id)
	for index, item := range travels {
		items[index] = item.Item
	}
	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetItemsByTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tripId := vars["tripId"]
	Id, err := strconv.ParseInt(tripId, 0, 0)
	if err != nil {
		fmt.Println("An error occured while parsing value")
	}

	items := models.GetItemsByTrip(Id)
	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	CreateItem := &models.Item{}
	utils.ParseBody(r, CreateItem)
	u := CreateItem.CreateItem()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	Id, err := strconv.ParseInt(itemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	item := models.DeleteItem(Id)
	res, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteItemsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	var items []*models.Item
	sends := models.DeleteItemsByUser(Id)
	for index, item := range sends {
		items[index] = item.Item
	}
	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteItemsByTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	Id, err := strconv.ParseInt(itemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	item := models.DeleteItemsByTrip(Id)
	res, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updateItem = &models.Item{}
	utils.ParseBody(r, updateItem)

	vars := mux.Vars(r)
	itemId := vars["itemId"]
	Id, err := strconv.ParseInt(itemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	itemDetail, db := models.GetItem(Id)

	db.Save(&itemDetail)

	res, _ := json.Marshal(itemDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
