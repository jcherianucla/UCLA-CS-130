package middleware

import (
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
)

func Logging(
	w http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc) {
	utilities.Logger.Infof("Starting request")
	next(w, r)
	utilities.Logger.Infof("Sent response")
}
