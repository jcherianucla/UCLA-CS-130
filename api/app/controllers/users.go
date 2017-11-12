package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"net/http"
)

var UsersShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
	},
)

var UsersBOL = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
	},
)

var UsersLogin = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		user, err := models.NewUser(r)
		var status int
		var msg, token string
		if err != nil {
			status = 500
			msg = err.Error()
		} else {
			status = 200
			msg = fmt.Sprintf("Created user with email: %s", user.Email)
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
