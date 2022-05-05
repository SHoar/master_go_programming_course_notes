package main

import (
	"log"
	"net/http"

	"github.com/SHoar/master_go_programming_course/scratchpad/goProductApi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Create new router
	router := mux.NewRouter()

	// route properly to respective handlers
	router.Handle("/products", handlers.CreateProductHandler()).Methods("POST")
	router.Handle("/products", handlers.GetProductHandler()).Methods("GET")

	// Create a server and assign the router
	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	log.Println("Starting Product Catalog server on Port 9090")
	server.ListenAndServe()

}
