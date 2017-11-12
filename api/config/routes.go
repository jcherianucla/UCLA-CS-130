package config

import (
	"github.com/gorilla/mux"
	"github.com/jcherianucla/UCLA-CS-130/api/app/controllers"
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	URI     string
	Handler http.HandlerFunc
}

type Routes struct {
	userRoutes    []Route
	classRoutes   []Route
	projectRoutes []Route
}

func (routes *Routes) createUserRoutes() {
	routes.userRoutes = Routes{
		Route{
			"Index",
			"GET",
			GetAPIInstance.Gen("/"),
			controllers.UserIndex,
		},
		Route{
			"Login",
			"POST",
			GetAPIInstance.Gen("/login"),
			middleware.Authenticate(controllers.UserLogin),
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
			Handler(route.Handler)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Holds all our routes
	routes := &Routes{}
	// Create the user routes
	routes.createUserRoutes()
	BindRoutes(&router, routes.userRoutes)
}
