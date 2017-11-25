package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"net/http"
	"strconv"
)

var ClassesIndex = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var status int
		var msg string
		strId := getClaims(r)
		creator_id, _ := strconv.Atoi(strId)
		classes, err := models.LayerInstance().Class.Get(models.ClassQuery{Creator_Id: creator_id}, "")
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
		class, err := models.NewClass(r)
		creator_id := getClaims(r)
		class, err = models.LayerInstance().Class.Insert(class, creator_id)
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
		class, _ := models.LayerInstance().Class.GetByID(params["id"])
		if fmt.Sprintf("%v", class.Creator_Id) != getClaims(r) {
			status = 500
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
		class, err := models.LayerInstance().Class.GetByID(params["id"])
		if fmt.Sprintf("%v", class.Creator_Id) != getClaims(r) {
			status = 500
			msg = "Invalid permissions to delete class"
		} else {
			err = models.LayerInstance().Class.Delete(params["id"])
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
