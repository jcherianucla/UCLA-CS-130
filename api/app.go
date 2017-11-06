package main

import (
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/controllers"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Basic example of adding a route
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/classes", controllers.AllClasses).Methods("GET")

	// Setup default middleware
	n := negroni.Classic()
	n.UseHandler(r)

	// Start server on specified port
	http.ListenAndServe(utilities.PORT, n)
}
