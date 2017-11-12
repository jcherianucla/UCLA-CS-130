package router

import (
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/controllers"
	"github.com/jcherianucla/UCLA-CS-130/api/middleware"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	URI         string
	HandlerFunc http.Handler
}

type Routes struct {
	userRoutes    []Route
	classRoutes   []Route
	projectRoutes []Route
}

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
			"Professor Login",
			"POST",
			utilities.GetAPIInstance().Gen("/login"),
			controllers.UsersLogin,
		},
	}
}

/**
 * Bind all specified routes to the mux router
 */
func BindRoutes(r *mux.Router, routes []Route) {
	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.URI).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Holds all our routes
	routes := &Routes{}
	// Create the user routes
	routes.createUserRoutes()
	BindRoutes(router, routes.userRoutes)
	return router
}
