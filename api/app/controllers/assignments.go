package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
	"strconv"
)

// AssignmentsIndex provides the route for showcasing all the projects for the specified class.
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

// AssignmentsShow forks in three potential avenues based on the user's role and the deadline.
// If the user is a professor, they are provided with analytics for the given assignment based
// on all student submissions. If the user is a student and hit the endpoint before the deadline,
// then they are given the assignment details, else they are shown the results of their submission.
var AssignmentsShow = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		utilities.SetupResponse(&w)
		if r.Method != "OPTIONS" {
			params := mux.Vars(r)
			var status int
			msg := "Success"
			resultKey := "result"
			var result interface{}
			userId := getClaims(r)
			assign_id, _ := strconv.ParseInt(params["aid"], 10, 64)
			user, err := models.LayerInstance().User.GetByID(userId)
			assignment, err := models.LayerInstance().Assignment.GetByID(params["aid"])
			// Student
			if err == nil && !user.Is_professor {
				// Past deadline, showcase submission results
				if !utilities.BeforeDeadline(assignment.Deadline) {
					r, err := models.LayerInstance().Submission.GetByID(userId, params["aid"])
					if err == nil && len(assignment.Grade_script) > 0 {
						// Execute the grade script on the assignment to get back all the results
						s, res, err := models.Exec(assignment.Grade_script, r.File, assignment.Lang)
						if err == nil {
							r.Score = s
							r.Post_results = res
							resultKey = "submission"
							temp := make(map[string]interface{})
							temp["submission"] = r
							temp["grade_script"] = string(assignment.Grade_script)
							result = temp
						}
					} else {
						msg = "You have no submission"
					}
				}
			} else {
				// Professor receives analytics
				resultKey = "analytics"
				r := make(map[string]interface{})
				students, err := models.LayerInstance().Enrolled.GetStudents(params["cid"])
				submissions, err := models.LayerInstance().Submission.Get(models.SubmissionQuery{Assignment_id: assign_id}, "")
				utilities.Sugar.Infof("Number of submissions: %v", len(submissions))
				utilities.Sugar.Infof("Number of students: %v", len(students))
				if err == nil {
					// Gets back total submisison ratio and the score breakdown
					if len(students) > 0 {
						r["num_submissions"] = float64(len(submissions)) / float64(len(students))
						var s []int64
						for _, submission := range submissions {
							s = append(s, submission.Score)
						}
						r["score"] = s
						result = r
					}
				}
			}
			if err != nil {
				status = 500
				msg = err.Error()
			} else {
				status = 200
			}
			if result == nil {
				result = ""
			}
			JSON, _ := json.Marshal(map[string]interface{}{
				"status":     status,
				"message":    msg,
				"assignment": assignment,
				resultKey:    result,
			})
			w.Write(JSON)
		}
	},
)

// AssignmentsCreate allows a professor to create an assignments,
// providing a grading and sanity script.
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

// AssignmentsUpdate allows a professor to update the assignment.
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

// AssignmentsDelete allows a professor to delete the assignment.
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
