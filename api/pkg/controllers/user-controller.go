package controllers

import (
	"api/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var NewUser models.User

func CreateUser(w http.ResponseWriter, r http.Request) {

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := uuid.Parse(userId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
