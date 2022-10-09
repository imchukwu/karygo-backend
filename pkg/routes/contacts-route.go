package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/controllers"
)

var RegisterContactRoutes = func(r *mux.Router) {
	r.HandleFunc("/contacts", controllers.CreateContact).Methods("POST")
	r.HandleFunc("/contacts", controllers.GetContacts).Methods("GET")
	r.HandleFunc("/contacts/{contactId}", controllers.GetContact).Methods("GET")
	r.HandleFunc("/contacts/{contactId}", controllers.UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{contactId}", controllers.DeleteContact).Methods("DELETE")
}
