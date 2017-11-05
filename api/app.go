package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	mux := mux.NewServeMux()
	// Basic example of adding a route
	mux.HandleFunc("/classes", controllers.AllClasses).Methods("GET")

	// Setup default middleware
	n := negroni.Classic()
	n.UseHandler(mux)

	// Start server on specified port
	http.ListenAndServe(utilities.PORT, n)
}
