package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/imchukwu/karygo_backend/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterItemRoutes(r)
	routes.RegisterTripRoutes(r)
	routes.RegisterContactRoutes(r)
	routes.RegisterAdminRoutes(r)
	routes.RegisterBillingRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Starting server %v", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
