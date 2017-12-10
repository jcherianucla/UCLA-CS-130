package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
	"strconv"
)

var AssignmentsIndex = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			var status int
			var msg string
			params := mux.Vars(r)
			class_id, _ := strconv.ParseInt(params["cid"], 10, 64)
			assignments, err := models.LayerInstance().Assignment.Get(models.AssignmentQuery{Class_id: class_id}, "")
			if err != nil {
				status = 500
				msg = err.Error()
			} else {
				status = 200
				msg = "Success"
			}
			JSON, _ := json.Marshal(map[string]interface{}{
				"status":      status,
				"message":     msg,
				"assignments": assignments,
			})
			w.Write(JSON)
		}
	},
)

var AssignmentsShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			params := mux.Vars(r)
			var status int
			var msg string
			var result interface{}
			userId := getClaims(r)
			user, err := models.LayerInstance().User.GetByID(userId)
			if !user.Is_professor {
				result, err = models.LayerInstance().Assignment.GetByID(params["aid"])
				if !utilities.BeforeDeadline(result.(models.Assignment).Deadline) {

				}
			}
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
				"result":  result,
			})
			w.Write(JSON)
		}
	},
)

var AssignmentsCreate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			var status int
			var msg string
			creator_id := getClaims(r)
			params := mux.Vars(r)
			assignment, err := models.NewAssignment(r)
			user, err := models.LayerInstance().User.GetByID(creator_id)
			if err != nil || !user.Is_professor {
				status = 400
				msg = "Invalid permissions to create an assignment"
			} else {
				class_id, _ := strconv.ParseInt(params["cid"], 10, 64)
				_, err := models.LayerInstance().Class.GetByID(params["cid"])
				if err == nil {
					assignment.Class_id = class_id
					assignment, err = models.LayerInstance().Assignment.Insert(assignment)
					if err == nil {
						status = 200
						msg = "Success"
					}
				}
				if err != nil {
					status = 500
					msg = err.Error()
				}
			}
			JSON, _ := json.Marshal(map[string]interface{}{
				"status":     status,
				"message":    msg,
				"assignment": assignment,
			})
			w.Write(JSON)
		}
	},
)

var AssignmentsUpdate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			params := mux.Vars(r)
			var status int
			var msg string
			creator_id := getClaims(r)
			updated, err := models.LayerInstance().Assignment.GetByID(params["aid"])
			if err != nil || !hasPermissions(creator_id, strconv.FormatInt(updated.Class_id, 10)) {
				status = 400
				msg = "Invalid permissions to update assignment"
			} else {
				updates, err := models.NewAssignment(r)
				updated, err = models.LayerInstance().Assignment.Update(params["aid"], updates)
				if err != nil {
					status = 500
					msg = err.Error()
				} else {
					status = 200
					msg = "Success"
				}
			}
			JSON, _ := json.Marshal(map[string]interface{}{
				"status":     status,
				"message":    msg,
				"assignment": updated,
			})
			w.Write(JSON)
		}
	},
)

var AssignmentsDelete = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			var status int
			var msg string
			params := mux.Vars(r)
			creator_id := getClaims(r)
			assignment, err := models.LayerInstance().Assignment.GetByID(params["aid"])
			if !hasPermissions(creator_id, strconv.FormatInt(assignment.Class_id, 10)) {
				status = 500
				msg = "Invalid permissions to delete assignment"
			} else {
				err = models.LayerInstance().Assignment.Delete(params["aid"])
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
		}
	},
)
