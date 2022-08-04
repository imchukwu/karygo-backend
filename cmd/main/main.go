package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterItemRoutes(r)
	routes.RegisterTripRoutes(r)
	http.Handle("/", r)

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
