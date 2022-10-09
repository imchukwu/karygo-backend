package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/controllers"
)

var RegisterContactRoutes = func(r *mux.Router) {
	r.HandleFunc("/contact", controllers.CreateContact).Methods("POST")
	r.HandleFunc("/contact", controllers.GetContacts).Methods("GET")
	r.HandleFunc("/contact/{contactId}", controllers.GetContact).Methods("GET")
	r.HandleFunc("/contact/{contactId}", controllers.UpdateContact).Methods("PUT")
	r.HandleFunc("/contact/{contactId}", controllers.DeleteContact).Methods("DELETE")
}
