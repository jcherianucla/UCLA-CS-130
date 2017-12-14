package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
	"strconv"
	"strings"
)

// hasPermissions ensures that for the given class, the user has permissions to invoke certain actions.
// It takes in the user id and class id to run the analysis on.
// It returns whether the user is the creator and a professor.
func hasPermissions(creator_id, class_id string) bool {
	user, err := models.LayerInstance().User.GetByID(creator_id)
	class, err := models.LayerInstance().Class.GetByID(class_id)
	return err == nil &&
		fmt.Sprintf("%v", class.Creator_id) == creator_id &&
		user.Is_professor
}

// ClassesIndex is used to return all the classes for the specified user as found from the authorization token.
var ClassesIndex = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			var status int
			var msg string
			var classes []models.Class
			strId := getClaims(r)
			user, err := models.LayerInstance().User.GetByID(strId)
			// Differentiate between student and professor
			if user.Is_professor {
				creator_id, _ := strconv.ParseInt(strId, 10, 64)
				classes, err = models.LayerInstance().Class.Get(models.ClassQuery{Creator_id: creator_id}, "")
			} else {
				classes, err = models.LayerInstance().Enrolled.GetClasses(strId)
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
				"classes": classes,
			})
			w.Write(JSON)
		}
	},
)

// ClassesShow shows details about the speicfic class given the class id in the url parameters.
var ClassesShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			params := mux.Vars(r)
			var status int
			var msg string
			class, err := models.LayerInstance().Class.GetByID(params["id"])
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
				"class":   class,
			})
			w.Write(JSON)
		}
	},
)

// ClassesCreate creates a class if the user is a professor and the class doesn't already exist for that same professor. Auto-enrolls students into the class from a given csv.
var ClassesCreate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			var status int
			var msg string
			class, err := models.NewClass(r)
			creator_id := getClaims(r)
			user, err := models.LayerInstance().User.GetByID(creator_id)
			if err != nil || !user.Is_professor {
				status = 400
				msg = "Invalid permissions to create a class."
			} else {
				class.Creator_id, _ = strconv.ParseInt(creator_id, 10, 64)
				class, err = models.LayerInstance().Class.Insert(class)
				if err != nil {
					status = 500
					msg = err.Error()
				} else {
					// Enroll students from csv
					f, _, err := r.FormFile("myfile")
					if err == nil {
						err = models.LayerInstance().Enrolled.Insert(strconv.FormatInt(class.Id, 10), f)
					}
					if f != nil && err != nil {
						_ = models.LayerInstance().Class.Delete(strconv.FormatInt(class.Id, 10))
						status = 500
						msg = err.Error()
					} else {
						status = 200
						msg = "Success"
					}
				}
			}
			JSON, _ := json.Marshal(map[string]interface{}{
				"status":  status,
				"message": msg,
				"class":   class,
			})
			w.Write(JSON)
		}
	},
)

// ClassesUpdate updates the specific class if the user is a professor and the creator of the class, and allows to enroll more students based on a newly submitted csv if one exists.
var ClassesUpdate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			params := mux.Vars(r)
			var status int
			var msg string
			var updated = models.Class{}
			creator_id := getClaims(r)
			if !hasPermissions(creator_id, params["id"]) {
				status = 400
				msg = "Invalid permissions to update class"
			} else {
				updates, err := models.NewClass(r)
				updated, err = models.LayerInstance().Class.Update(params["id"], updates)
				// Enroll new students
				f, _, err := r.FormFile("myfile")
				if err == nil {
					err = models.LayerInstance().Enrolled.Insert(params["id"], f)
				}
				if err != nil && !strings.Contains(err.Error(), "no such file") {
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
				"class":   updated,
			})
			w.Write(JSON)
		}
	},
)

// ClassesDelete will provide the route to delete the specified class if the user has permissions.
var ClassesDelete = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			params := mux.Vars(r)
			var status int
			var msg string
			creator_id := getClaims(r)
			if !hasPermissions(creator_id, params["id"]) {
				status = 500
				msg = "Invalid permissions to delete class"
			} else {
				// Delete class and the enrollments
				err := models.LayerInstance().Class.Delete(params["id"])
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
