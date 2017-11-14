// The router package acts as the intermediary between the frontend
// service and the backend, using the controller to provide the handlers
// for each route.
package router

import (
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/controllers"
	"github.com/jcherianucla/UCLA-CS-130/api/middleware"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
)

// Represents all components of a route. Convenient
// use with mux router.
type Route struct {
	Name        string
	Method      string
	URI         string
	HandlerFunc http.Handler
}

// Represents all the routes for the application.
type Routes struct {
	userRoutes    []Route
	classRoutes   []Route
	projectRoutes []Route
}

// createUserRoutes will instantiate all routes for
// users.
func (routes *Routes) createUserRoutes() {
	routes.userRoutes = []Route{
		Route{
			"Show",
			"GET",
			utilities.GetAPIInstance().Gen("/users/{id}"),
			middleware.Authenticate(controllers.UsersShow),
		},
		Route{
			"BOL",
			"POST",
			utilities.GetAPIInstance().Gen("/bol"),
			controllers.UsersBOL,
		},
		Route{
			"Create",
			"POST",
			utilities.GetAPIInstance().Gen("/users"),
			controllers.UsersCreate,
		},
		Route{
			"Professor Login",
			"POST",
			utilities.GetAPIInstance().Gen("/login"),
			controllers.UsersLogin,
		},
	}
}

// BindRoutes will bind all routes to the router.
// It takes in a reference to the router and the slice
// of all routes.
func BindRoutes(r *mux.Router, routes []Route) {
	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.URI).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

// NewRouter instantiates a new router with all routes
// bound to the router.
// It returns an instance to the router.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Holds all our routes
	routes := &Routes{}
	// Create the user routes
	routes.createUserRoutes()
	BindRoutes(router, routes.userRoutes)
	return router
}
