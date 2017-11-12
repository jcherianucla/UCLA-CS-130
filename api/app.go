package main

import (
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/config/router"
	"github.com/jcherianucla/UCLA-CS-130/api/middleware"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/urfave/negroni"
	"net/http"
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

	// Initialize global singleton with init
	_ = models.LayerInstance()
	// Spin up server
	http.ListenAndServe(utilities.PORT, n)

	// Create DB
	/*
		db, err := db.New(db.Config{
			utilities.DB_HOST,
			utilities.DB_PORT,
			utilities.DB_USER,
			utilities.DB_PASSWORD,
			utilities.DB_NAME,
		})
		utilities.CheckError(err)
	*/
	// Test
	//ut, err := models.NewUserTable(&db)
	/*
				_, err = ut.InsertUser(models.User{
					Is_professor: false,
					Email:        "jcherian@ucla.edu",
					First_name:   "Jahan",
					Last_name:    "Kuruvilla Cherian",
					Password:     []byte("swag"),
				})
				utilities.CheckError(err)

			_, err = ut.UpdateUser(2, models.User{
				Is_professor: true,
				Email:        "das@cs.ucla.edu",
				First_name:   "David",
				Last_name:    "Smallberg",
			})
			utilities.CheckError(err)
		err = ut.DeleteUser(2)
		utilities.CheckError(err)
	*/
}
