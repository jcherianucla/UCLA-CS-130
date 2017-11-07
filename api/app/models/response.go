package models

import "time"

type Response struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Requested  interface{} `json:"request"`
	Created_at time.Time   `json:"created_at"`
}

func CreateResponse(status int, msg string, req interface{}) *Response {
	resp := Response{status, msg, req, time.Now()}
	return &resp
}
