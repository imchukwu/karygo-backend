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

// swagger:route GET /users users listUsers
// 	Returns a list of users
// 	Responses:
//  	200: usersResponse
//  Produces:
// 		- application/json
//		- application/xml

// list of users to be returned in response
// swagger:response usersResponse
type usersResponse struct{
	// All users in the system
	// in: body
	Body []models.User

}

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	users := models.GetUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route GET /users/{userId} user getUser
// 	Returns a user by ID
// 	Parameters:
// 		name: userId
// 		in: path
// 		description: ID of user to be fetched
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: userResponse
//  Produces:
// 		- application/json
//		- application/xml

// user to be returned in response
// swagger:response userResponse
type userResponse struct{
	// A user in the system
	// in: body
	Body models.User

}


// Get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetail, _ := models.GetUser(Id)
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route POST /users user creatuser
// 	Creates new user
// 	responses:
//  	200: createUserResponse
// 		
// 		default: error
//  Produces:
// 		- application/json
//		- application/xml

// New user to be returned in response
// swagger:response createUserResponse
type createUserResponse struct{
	// Create new user in the system
	// in: body
	Body models.User

}

// Create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	u := CreateUser.CreateUser()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route DELETE /users/{userId} user deleteUser
// 	Returns deleted user
// 	Parameters:
// 		name: userId
// 		in: path
// 		description: ID of user to be deleted
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: deleteUserResponse
//  Produces:
// 		- application/json
//		- application/xml

// deleted user is returned in response
// swagger:response deleteUserResponse
type deleteUserResponse struct{
	// delete user in the system
	// in: body
	Body models.User

}




// Delete user from the system by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(Id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// swagger:route PUT /users/{userId} user updateUser 
// 	Returns updated user
// 	Parameters:
// 		name: userId
// 		in: path
// 		description: ID of user to be updated
// 		required: true
// 		type: integer
// 		format: int64
// 	responses:
//  	200: updateUserResponse
//  Produces:
// 		- application/json
//		- application/xml

// updated user is returned in response
// swagger:response updateUserResponse
type updateUserResponse struct{
	// update user in the system
	// in: body
	Body models.User

}


// update an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)

	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetail, db := models.GetUser(Id)
	if updateUser.Firstname != "" {
		userDetail.Firstname = updateUser.Firstname
	}
	if updateUser.Middlename != "" {
		userDetail.Middlename = updateUser.Middlename
	}
	if updateUser.Lastname != "" {
		userDetail.Lastname = updateUser.Lastname
	}
	if updateUser.Email != "" {
		userDetail.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		userDetail.Password = updateUser.Password
	}
	if updateUser.Address != "" {
		userDetail.Address = updateUser.Address
	}
	if updateUser.Gender != "" {
		userDetail.Gender = updateUser.Gender
	}
	if updateUser.RegDate != "" {
		userDetail.RegDate = updateUser.RegDate
	}
	// if updateUser.State.Name != "" {
	// 	userDetail.State.Name = updateUser.State.Name
	// }
	// if updateUser.State.Lga.Name != "" {
	// 	userDetail.State.Lga.Name = updateUser.State.Lga.Name
	// }

	//TODO: Add Trips

	db.Save(&userDetail)

	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
