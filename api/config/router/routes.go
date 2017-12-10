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
			"BOL",
			"POST",
			utilities.GetAPIInstance().Gen("/bol"),
			controllers.UsersBOL,
		},
		Route{
			"Professor Login",
			"POST",
			utilities.GetAPIInstance().Gen("/login"),
			controllers.UsersLogin,
		},
		Route{
			"Create",
			"POST",
			utilities.GetAPIInstance().Gen("/users"),
			controllers.UsersCreate,
		},
		Route{
			"Show Current",
			"GET",
			utilities.GetAPIInstance().Gen("/user"),
			middleware.Authenticate(controllers.UsersIndex),
		},
		Route{
			"Show User",
			"GET",
			utilities.GetAPIInstance().Gen("/users/{id}"),
			middleware.Authenticate(controllers.UsersShow),
		},
		Route{
			"Update",
			"PUT",
			utilities.GetAPIInstance().Gen("/users/{id}"),
			middleware.Authenticate(controllers.UsersUpdate),
		},
		Route{
			"Delete",
			"DELETE",
			utilities.GetAPIInstance().Gen("/users/{id}"),
			middleware.Authenticate(controllers.UsersDelete),
		},
	}
}

// createClassRoutes will instantiate all routes for
// classes.
func (routes *Routes) createClassRoutes() {
	routes.classRoutes = []Route{
		Route{
			"Create",
			"POST",
			utilities.GetAPIInstance().Gen("/classes"),
			middleware.Authenticate(controllers.ClassesCreate),
		},
		Route{
			"Show All",
			"GET",
			utilities.GetAPIInstance().Gen("/classes"),
			middleware.Authenticate(controllers.ClassesIndex),
		},
		Route{
			"Show",
			"GET",
			utilities.GetAPIInstance().Gen("/classes/{id}"),
			middleware.Authenticate(controllers.ClassesShow),
		},
		Route{
			"Update",
			"PUT",
			utilities.GetAPIInstance().Gen("/classes/{id}"),
			middleware.Authenticate(controllers.ClassesUpdate),
		},
		Route{
			"Delete",
			"DELETE",
			utilities.GetAPIInstance().Gen("/classes/{id}"),
			middleware.Authenticate(controllers.ClassesDelete),
		},
	}
}

// BindRoutes will bind all routes to the router.
// It takes in a reference to the router and the slice
// of all routes.
func BindRoutes(r *mux.Router, routes []Route) {
	for _, route := range routes {
		r.
			Methods(route.Method, "OPTIONS").
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
	// Create the class routes
	routes.createClassRoutes()
	BindRoutes(router, routes.userRoutes)
	BindRoutes(router, routes.classRoutes)
	return router
}
