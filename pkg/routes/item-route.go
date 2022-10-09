package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/controllers"
)

var RegisterItemRoutes = func(r *mux.Router) {
	r.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	r.HandleFunc("/items", controllers.GetItems).Methods("GET")
	r.HandleFunc("/items/{itemId}", controllers.GetItem).Methods("GET")
	r.HandleFunc("/items/{userId}/send", controllers.GetItemsSentByUser).Methods("GET")
	r.HandleFunc("/items/{userId}/travel", controllers.GetItemsTransportedByUser).Methods("GET")
	r.HandleFunc("/items/{itemId}", controllers.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{itemId}", controllers.DeleteItem).Methods("DELETE")
	// r.HandleFunc("/items/{userId}", controllers.DeleteItemsByUser).Methods("DELETE")
	// r.HandleFunc("/items/{tripId}", controllers.DeleteItemsByTrip).Methods("DELETE")
}
