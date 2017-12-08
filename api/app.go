// The main package spins up the overall backend service.
package main

import (
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/config/router"
	"github.com/jcherianucla/UCLA-CS-130/api/middleware"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	r := router.NewRouter()
	// Setup default middleware
	n := negroni.New(
		negroni.HandlerFunc(middleware.Logging),
		negroni.NewLogger(),
	)
	n.UseHandler(r)

	// Initialize global singleton with init
	_ = models.LayerInstance()
	// Spin up server
	http.ListenAndServe(":"+utilities.PORT, n)
}
