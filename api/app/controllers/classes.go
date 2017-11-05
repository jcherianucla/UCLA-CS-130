package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var AllClasses = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		// Define JSON response
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  200,
			"message": "Welcome to GradePortal",
		})
		w.Write(JSON)
	},
)
