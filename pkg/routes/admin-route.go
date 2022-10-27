package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/controllers"
)

var RegisterAdminRoutes = func(r *mux.Router) {
	r.HandleFunc("/admin", controllers.CreateAdmin).Methods("POST")
	r.HandleFunc("/admin", controllers.GetAdmins).Methods("GET")
	r.HandleFunc("/admin/{adminId}", controllers.GetAdmin).Methods("GET")
	r.HandleFunc("/admin/{adminId}", controllers.UpdateAdmin).Methods("PUT")
	r.HandleFunc("/admin/{admin}", controllers.DeleteAdmin).Methods("DELETE")
	r.HandleFunc("/admin/transactiion", controllers.GetTransactions).Methods("GET")
	r.HandleFunc("/admin/transaction/{txId}", controllers.GetTransaction).Methods("GET")
}
