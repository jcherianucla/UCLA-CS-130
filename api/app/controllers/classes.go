package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	// Testing
	//"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
	"strconv"
)

func hasPermissions(creator_id, class_id string) bool {
	user, err := models.LayerInstance().User.GetByID(creator_id)
	class, err := models.LayerInstance().Class.GetByID(class_id)
	return err == nil &&
		fmt.Sprintf("%v", class.Creator_id) == creator_id &&
		user.Is_professor
}

var ClassesIndex = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var status int
		var msg string
		strId := getClaims(r)
		creator_id, _ := strconv.ParseInt(strId, 10, 64)
		classes, err := models.LayerInstance().Class.Get(models.ClassQuery{Creator_id: creator_id}, "")
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
	},
)

var ClassesShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
	},
)

var ClassesCreate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
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
				if err != nil {
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
	},
)

var ClassesUpdate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
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
			"class":   updated,
		})
		w.Write(JSON)
	},
)

var ClassesDelete = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var status int
		var msg string
		creator_id := getClaims(r)
		if !hasPermissions(creator_id, params["id"]) {
			status = 500
			msg = "Invalid permissions to delete class"
		} else {
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
	},
)
