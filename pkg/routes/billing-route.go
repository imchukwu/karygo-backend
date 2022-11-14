package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/controllers"
)

var RegisterBillingRoutes = func (r *mux.Router){
	r.HandleFunc("/billing", controllers.CreateBilling).Methods("POST")
	r.HandleFunc("/billings", controllers.GetBillings).Methods("GET")
	r.HandleFunc("/billing/{billingId}", controllers.GetBilling).Methods("GET")
	r.HandleFunc("/billing/{billingId}", controllers.UpdateBilling).Methods("PUT")
	r.HandleFunc("/billing/{billingId}", controllers.DeleteBilling).Methods("DELETE")
	r.HandleFunc("/billing/{userId}/user", controllers.GetBillingsByUser).Methods("GET")
}