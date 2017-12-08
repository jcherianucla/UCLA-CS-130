package controllers

/*
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"net/http"
	"strconv"
)

var AssignmentsIndex = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var status int
		var msg string
		strId := getClaims(r)
		user_id, _ := strconv.ParseInt(strId, 10, 64)
		classes, err := models.LayerInstance().Assignment.Get(models.ClassQuery{Creator_id: creator_id}, "")
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

var AssignmentsShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var status int
		var msg string
		class, err := models.LayerInstance().Assignment.GetByID(params["id"])
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

var AssignmentsCreate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		var status int
		var msg string
		class, err := models.Ne.Assignment.r)
		creator_id := getClaims(r)
		user, err := models.LayerInstance().User.GetByID(creator_id)
		if err != nil || !user.Is_professor {
			status = 400
			msg = "Invalid permissions to create a class."
		} else {
			class, err = models.LayerInstance().Assignment.Insert(class, creator_id)
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
			"class":   class,
		})
		w.Write(JSON)
	},
)

var AssignmentsUpdate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var status int
		var msg string
		var updated = models.Assignment.}
		class, _ := models.LayerInstance().Assignment.GetByID(params["id"])
		creator_id := getClaims(r)
		user, err := models.LayerInstance().User.GetByID(creator_id)
		if err != nil || fmt.Sprintf("%v", class.Creator_id) != creator_id || !user.Is_professor {
			status = 400
			msg = "Invalid permissions to update class"
		} else {
			updates, err := models.Ne.Assignment.r)
			updated, err = models.LayerInstance().Assignment.Update(params["id"], updates)
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

var AssignmentsDelete = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var status int
		var msg string
		class, err := models.LayerInstance().Assignment.GetByID(params["id"])
		if fmt.Sprintf("%v", class.Creator_id) != getClaims(r) {
			status = 500
			msg = "Invalid permissions to delete class"
		} else {
			err = models.LayerInstance().Assignment.Delete(params["id"])
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
*/
