package routes

import (
	"github.com/dzoxploit/crud-contact-golang/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/contact/", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts", controllers.GetContact).Methods("GET")
	router.HandleFunc("/contact/{contactId}", controllers.GetContactById).Methods("GET")
	router.HandleFunc("/contact/{contactId}", controllers.UpdateContact).Methods("PUT")
	router.HandleFunc("/contact/{contactId}", controllers.DeleteContact).Methods("DELETE")
}