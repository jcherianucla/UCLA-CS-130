package controllers

import (
	"fmt"
	"net/http"
)

var Home = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	},
)
