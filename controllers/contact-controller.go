package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dzoxploit/crud-contact-golang/models"
	"github.com/dzoxploit/crud-contact-golang/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var NewContact models.Contact

func CreateContact(w http.ResponseWriter, r *http.Request) {

	CreateContact := &models.Contact{}
	utils.ParseBody(r, CreateContact)
	CreateContact.ID = uuid.New().String()
	b:= CreateContact.CreateContact()
	fmt.Print(b)
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	// Extract the page number and page size from query parameters
	pageNumber, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	pageSearch := r.URL.Query().Get("search")

	// Call the model function to get contacts with pagination
	contacts, err := models.GetAllContact(pageNumber, pageSize, pageSearch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the contacts to JSON
	response, err := json.Marshal(contacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetContactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactId := vars["contactId"]
	if len(contactId) != 0 {
		fmt.Println("Error Null")
	}
	contactDetails, _:= models.GetContactById(contactId)
	res, _ := json.Marshal(contactDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	var updateContact = &models.Contact{}
	utils.ParseBody(r, updateContact)
	vars := mux.Vars(r)
	contactId := vars["contactId"]
	if len(contactId) != 0 {
		fmt.Println("Error Null")
	}
	contactDetails, db:= models.GetContactById(contactId)
	if updateContact.Name != "" {
		contactDetails.Name = updateContact.Name
	}
	if updateContact.Email != "" {
		contactDetails.Email = updateContact.Email
	}
	if updateContact.Gender != "" {
		contactDetails.Gender = updateContact.Gender
	}
	if updateContact.Phone != "" {
		contactDetails.Phone = updateContact.Phone
	}
	db.Save(&contactDetails)
	res, _ := json.Marshal(contactDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactId := vars["contactId"]
	if len(contactId) != 0 {
		fmt.Println("Error null")
	}
	contact:= models.DeleteContact(contactId)
	res, _ := json.Marshal(contact)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}