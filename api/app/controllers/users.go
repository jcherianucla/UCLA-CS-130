// The controllers package house the Controller layer within the MVC architecture design.
// This is the middle layer in the architecture, which communicates from the view to the
// model. It does this through a RESTful API, as the View layer is a separate service.
// In effect the controller represents all the handlers exposed to the router upstream.
// All handlers take in the incoming HTTP request and a writer to the HTTP response,
// writing as JSON.
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"net/http"
)

// UsersShow will show the home page for the user.
var UsersShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	},
)

// UsersBOL will use Google Omniauth callback from the frontend
// to automatically insert a new student object.
var UsersBOL = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
	},
)

// UsersCreate will explicitly insert a user object.
// Used to create a professor.
var UsersCreate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		user, err := models.NewUser(r)
		fmt.Printf("%v", user)
		user, err = models.LayerInstance().User.Insert(user)
		var status int
		var msg string
		if err != nil {
			status = 500
			msg = err.Error()
		} else {
			status = 200
			msg = "Success"
		}
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  status,
			"message": msg,
			"user":    user,
		})
		w.Write(JSON)
	},
)

// UsersLogin will explicitly login a user, handing back
// a JWT for continued access to restricted parts of the site.
var UsersLogin = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		user, err := models.NewUser(r)
		user, err = models.LayerInstance().User.Login(user)
		var status int
		var msg, token string
		if err != nil {
			status = 500
			msg = err.Error()
		} else {
			status = 200
			msg = "Success"
			token = user.GenerateJWT()
		}
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  status,
			"message": msg,
			"token":   token,
		})
		w.Write(JSON)
	},
)
