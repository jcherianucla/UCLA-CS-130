package main

import (
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/config/router"
	"github.com/jcherianucla/UCLA-CS-130/api/middleware"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/urfave/negroni"
	//"net/http"
)

func main() {
	r := router.NewRouter()
	// Setup default middleware
	n := negroni.New(
		negroni.HandlerFunc(middleware.Logging),
		negroni.NewLogger(),
	)
	n.UseHandler(r)

	// Start server on specified port
	//http.ListenAndServe(utilities.PORT, n)

	// Create DB
	db, err := db.New(db.Config{
		utilities.DB_HOST,
		utilities.DB_PORT,
		utilities.DB_USER,
		utilities.DB_PASSWORD,
		utilities.DB_NAME,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Testing")
	// Test
	ut, err := models.NewUserTable(&db)
	_, err = ut.InsertUser(models.User{
		Id:           1,
		Is_professor: false,
		Email:        "jcherian@ucla.edu",
		First_name:   "Jahan",
		Last_name:    "Kuruvilla Cherian",
		Password:     []byte("swag"),
	})
	if err != nil {
		panic(err)
	}
}
