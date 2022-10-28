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

// swagger:route GET /contact contacts listContacts
// 	Returns a list of contacts
// 	Responses:
//  	200: contactsResponse
//  Produces:
// 		- application/json
//		- application/xml

// list of contacts to be returned in response
// swagger:response contactsResponse
type contactsResponse struct{
	// All contacts in the system
	// in: body
	Body []models.Contact

}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	contacts := models.GetContacts()
	res, _ := json.Marshal(contacts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route GET /contact/{contactId} contact getContact
// 	Returns an contact by ID
// 	Parameters:
// 		name: contactId
// 		in: path
// 		description: ID of contact to be fetched
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: contactResponse
//  Produces:
// 		- application/json
//		- application/xml

// contact to be returned in response
// swagger:response contactResponse
type contactResponse struct{
	// All contacts in the system
	// in: body
	Body models.Contact

}

// Get a contact by ID
func GetContact(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	contactId := vars["contactId"]
	Id, err := strconv.ParseInt(contactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	contactDetail, _ := models.GetContact(Id)
	res, _ := json.Marshal(contactDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  swagger:route POST /contact contact creatContact
// 	Creates new contact
// 	responses:
//  	200: createContactResponse
// 		
// 		default: error
//  Produces:
// 		- application/json
//		- application/xml

// New contact to be returned in response
// swagger:response createContactResponse
type createContactResponse struct{
	// Create new contact in the system
	// in: body
	Body models.Contact

}


// Creates new contact
func CreateContact(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	CreateContact := &models.Contact{}
	utils.ParseBody(r, CreateContact)
	contact := CreateContact.CreateContact()
	res, _ := json.Marshal(contact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route DELETE /contact/{contactId} contact deleteContact
// 	Returns deleted contact
// 	Parameters:
// 		name: contactId
// 		in: path
// 		description: ID of contact to be deleted
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: deleteContactResponse
//  Produces:
// 		- application/json
//		- application/xml

// deleted contact is returned in response
// swagger:response deleteContactResponse
type deleteContactResponse struct{
	// delete contact in the system
	// in: body
	Body models.Contact

}


// Delete contact
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	contactId := vars["contactId"]
	Id, err := strconv.ParseInt(contactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	contact := models.DeleteContact(Id)
	res, _ := json.Marshal(contact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// swagger:route PUT /contact/{contactId} contact updateContact 
// 	Returns updated contact
// 	Parameters:
// 		name: contactId
// 		in: path
// 		description: ID of contact to be updated
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: updateContactResponse
//  Produces:
// 		- application/json
//		- application/xml

// updated contact is returned in response
// swagger:response updateContactResponse
type updateContactResponse struct{
	// update contact in the system
	// in: body
	Body models.Contact

}


// Updates a contact by ID
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var updateContact = &models.Contact{}
	utils.ParseBody(r, updateContact)

	vars := mux.Vars(r)
	contactId := vars["contactId"]
	Id, err := strconv.ParseInt(contactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	contactDetail, db := models.GetContact(Id)
	if updateContact.Name != "" {
		contactDetail.Name = updateContact.Name
	}
	if updateContact.Email != "" {
		contactDetail.Email = updateContact.Email
	}
	if updateContact.Phone != "" {
		contactDetail.Phone = updateContact.Phone
	}
	if updateContact.Message != "" {
		contactDetail.Message = updateContact.Message
	}

	db.Save(&contactDetail)

	res, _ := json.Marshal(contactDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
