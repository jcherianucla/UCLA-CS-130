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
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
	"strings"
)

// getClaims will extract the authorization token from a request and get the associated claims for that id.
func getClaims(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")[len("Bearer "):]
	claims := utilities.ExtractClaims(tokenString)
	return fmt.Sprintf("%v", claims["id"])
}

// UsersIndex will show the home page for the user based
// purely on the claims.
var UsersIndex = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		user_id := getClaims(r)
		var status int
		var msg string
		user, err := models.LayerInstance().User.GetByID(user_id)
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

// UsersShow will show the home page for a user based
// a provided id as a parameter.
var UsersShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var status int
		var msg string
		user, err := models.LayerInstance().User.GetByID(params["id"])
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

// UsersBOL will use Google Omniauth callback from the frontend
// to automatically insert a new student object.
var UsersBOL = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		user, err := models.NewUser(r)
		// BOL is for students only
		user.Is_professor = false
		user.Password = utilities.DEFAULT_PASSWORD
		var status int
		var msg, token string
		_, err = models.LayerInstance().User.Insert(user)
		if err == nil || strings.Contains(err.Error(), "User already exists") {
			user, err = models.LayerInstance().User.Login(user)
			if err != nil {
				status = 500
				msg = err.Error()
			} else {
				status = 200
				msg = "Success"
				token = user.GenerateJWT()
			}
		} else {
			status = 500
			msg = err.Error()
		}
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  status,
			"message": msg,
			"token":   token,
		})
		w.Write(JSON)
	},
)

// UsersCreate will explicitly insert a user object.
// Used to create a professor.
var UsersCreate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		user, err := models.NewUser(r)
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

var UsersUpdate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var status int
		var msg string
		var updated = models.User{}
		if params["id"] != getClaims(r) {
			status = 500
			msg = "Invalid permissions to update user"
		} else {
			updates, err := models.NewUser(r)
			updated, err = models.LayerInstance().User.Update(params["id"], updates)
			if err != nil {
				status = 500
				msg = err.Error()
			} else {
				status = 200
				msg = "Success"
			}
		}
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  status,
			"message": msg,
			"user":    updated,
		})
		w.Write(JSON)
	},
)

var UsersDelete = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var status int
		var msg string
		if params["id"] != getClaims(r) {
			status = 500
			msg = "Invalid permissions to delete user"
		} else {
			err := models.LayerInstance().User.Delete(params["id"])
			if err != nil {
				status = 500
				msg = err.Error()
			} else {
				status = 200
				msg = "Success"
			}
		}
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  status,
			"message": msg,
		})
		w.Write(JSON)
	},
)
