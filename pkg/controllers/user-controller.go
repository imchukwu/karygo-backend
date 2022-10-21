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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	users := models.GetUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

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
