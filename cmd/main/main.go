package main

import (
	"log"
	"net/http"

	"github.com/ascendantaditya/Go-BookManagement/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	// This is the main function
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", nil))
}
