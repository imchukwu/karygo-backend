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

// swagger:route GET /items items listItems
// 	Returns a list of items
// 	Responses:
//  	200: itemsResponse
//  Produces:
// 		- application/json
//		- application/xml

// list of items to be returned in response
// swagger:response itemsResponse
type itemsResponse struct{
	// All items in the system
	// in: body
	Body []models.Item

}
// Get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	items := models.GetItems()
	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route GET /items/{itemId} item getItem
// 	Returns an item by ID
// 	Parameters:
// 		name: itemId
// 		in: path
// 		description: ID of item to be fetched
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: itemResponse
//  Produces:
// 		- application/json
//		- application/xml

// item to be returned in response
// swagger:response itemResponse
type itemResponse struct{
	// All items in the system
	// in: body
	Body models.Item

}


// Get item by ID
func GetItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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


// swagger:route GET /items/{userId}/send item getItemsSentByUSer
// 	Returns list of items by userId
// 	Parameters:
// 		name: userId
// 		in: path
// 		description: items sent by the userId
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: userSentItemsResponse
//  Produces:
// 		- application/json
//		- application/xml

// items to be returned in response
// swagger:response userSentItemsResponse
type userSentItemsResponse struct{
	// All items in the system
	// in: body
	Body []models.Item

}


// Items sent by user
func GetItemsSentByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	var itemIds []uint
	var items []*models.Item
	itemIds = models.GetItemIdBySender(Id)
	for index := range itemIds {
		itemDetail, _ := models.GetItem(Id)
		items[index] = itemDetail
	}

	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route GET /items/{userId}/travel item getItemsTransportedByUSer
// 	Returns list of items by userId
// 	Parameters:
// 		name: userId
// 		in: path
// 		description: items sent by the userId
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: userTransportedItemsResponse
//  Produces:
// 		- application/json
//		- application/xml

// items to be returned in response
// swagger:response userTransportedItemsResponse
type userTransportedItemsResponse struct{
	// All items in the system
	// in: body
	Body []models.Item

}



// Items transported by user
func GetItemsTransportedByUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("An error ocurred while parsing value")
	}

	var itemIds []uint
	var items []*models.Item
	itemIds = models.GetItemIdByTraveler(Id)
	for index := range itemIds {
		itemDetail, _ := models.GetItem(Id)
		items[index] = itemDetail
	}

	res, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route POST /items item creatItem
// 	Creates new item
// 	responses:
//  	200: createItemResponse
// 		
// 		default: error
//  Produces:
// 		- application/json
//		- application/xml

// New item to be returned in response
// swagger:response createItemResponse
type createItemResponse struct{
	// Create new item in the system
	// in: body
	Body models.Item

}



// Add new item to the system
func CreateItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	CreateItem := &models.Item{}
	utils.ParseBody(r, CreateItem)
	u := CreateItem.CreateItem()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route DELETE /items/{itemId} item deleteitem
// 	Returns deleted item
// 	Parameters:
// 		name: itemId
// 		in: path
// 		description: ID of item to be deleted
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: deleteItemResponse
//  Produces:
// 		- application/json
//		- application/xml

// deleted item is returned in response
// swagger:response deleteItemResponse
type deleteItemResponse struct{
	// delete item in the system
	// in: body
	Body models.Item

}


// Delete item from the system
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

// func DeleteItemsByUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userId := vars["userId"]
// 	Id, err := strconv.ParseInt(userId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}

// 	var items []*models.Item
// 	sends := models.DeleteItemsByUser(Id)
// 	for index, item := range sends {
// 		items[index] = item.Item
// 	}
// 	res, _ := json.Marshal(items)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteItemsByTrip(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	itemId := vars["itemId"]
// 	Id, err := strconv.ParseInt(itemId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	item := models.DeleteItemsByTrip(Id)
// 	res, _ := json.Marshal(item)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// swagger:route PUT /item/{itemId} item updateItem 
// 	Returns updated item
// 	Parameters:
// 		name: itemId
// 		in: path
// 		description: ID of item to be updated
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: updateItemResponse
//  Produces:
// 		- application/json
//		- application/xml

// updated item is returned in response
// swagger:response updateItemResponse
type updateItemResponse struct{
	// update item in the system
	// in: body
	Body models.Item

}


// Updates the existing item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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
