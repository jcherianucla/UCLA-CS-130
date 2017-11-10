package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/controllers"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/config"
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

	// Create DB
	db, err := config.New(config.Config{
		utilities.DB_HOST,
		utilities.DB_PORT,
		utilities.DB_USER,
		utilities.DB_PASSWORD,
		utilities.DB_NAME,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// Test
	ut, err := models.NewUserTable(&db)
	ut.InsertUser(models.User{
		Id:           1,
		Is_professor: false,
		Email:        "jcherian@ucla.edu",
		First_name:   "Jahan",
		Last_name:    "Kuruvilla Cherian",
		Password:     []byte("swag"),
	})
}
