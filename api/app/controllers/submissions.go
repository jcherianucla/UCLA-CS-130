package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
	"strconv"
	"time"
)

var SubmissionsCreate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			var status int
			msg := "Success"
			userId := getClaims(r)
			params := mux.Vars(r)
			submission, err := models.NewSubmission(r)
			user, err := models.LayerInstance().User.GetByID(userId)
			if !user.Is_professor {
				uid, _ := strconv.ParseInt(userId, 10, 64)
				submission.User_id = uid
				assign_id, _ := strconv.ParseInt(params["aid"], 10, 64)
				submission.Assignment_id = assign_id
				assignment, err := models.LayerInstance().Assignment.GetByID(params["aid"])
				submission.Time_updated = time.Now()
				submission, err = models.LayerInstance().Submission.Insert(submission)
				if err == nil && utilities.BeforeDeadline(assignment.Deadline) {
					if len(assignment.Sanity_script) > 0 {
						_, res, err := models.Exec(assignment.Sanity_script, submission.File, assignment.Lang)
						if err == nil {
							submission.Pre_results = res
						}
					}
				}
				if err != nil {
					status = 500
					msg = err.Error()
				}
			} else {
				status = 400
				msg = "You cannot submit a project"
			}
			if err != nil {
				status = 500
				msg = err.Error()
			} else {
				if status == 0 {
					status = 200
				}
			}
			JSON, _ := json.Marshal(map[string]interface{}{
				"status":     status,
				"message":    msg,
				"submission": submission,
			})
			w.Write(JSON)
		}
	},
)

var SubmissionsUpdate = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			params := mux.Vars(r)
			var status int
			msg := "Success"
			userId := getClaims(r)
			updated, err := models.LayerInstance().Submission.GetByID(userId, params["aid"])
			assignment, err := models.LayerInstance().Assignment.GetByID(params["aid"])
			updates, err := models.NewSubmission(r)
			sid := strconv.FormatInt(updated.Id, 10)
			updated.Time_updated = time.Now()
			updated, _ = models.LayerInstance().Submission.Update(sid, updates)
			if utilities.BeforeDeadline(assignment.Deadline) {
				if len(assignment.Sanity_script) > 0 {
					_, res, err := models.Exec(assignment.Sanity_script, updated.File, assignment.Lang)
					if err == nil {
						updated.Pre_results = res
					}
				}
			} else {
				status = 400
				msg = "Cannot update submission after deadline"
			}
			if err != nil {
				if status == 0 {
					status = 500
					msg = err.Error()
				}
			} else {
				if status == 0 {
					status = 200
				}
			}
			JSON, _ := json.Marshal(map[string]interface{}{
				"status":     status,
				"message":    msg,
				"submission": updated,
			})
			w.Write(JSON)
		}
	},
)
