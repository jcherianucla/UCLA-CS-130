// The utilities package provides additional functionality
// available to all packages within the application.
package utilities

import (
	"fmt"
	"net/http"
	"sync"
)

// Represents an api version for future
// extensibility.
type api struct {
	base int
	sub  int
}

// UpVersion increases the version number for the current
// api instance.
func (api *api) UpVersion() {
	if api.sub < 9 {
		api.sub += 1
	} else {
		api.base += 1
		api.sub = 0
	}
}

// Gen creates a prepended string for api versioned routes.
// It takes in a string to prepend to.
// It returns the prepended string.
func (api *api) Gen(uri string) string {
	return fmt.Sprintf("/api/v%d.%d%s", api.base, api.sub, uri)
}

var instance *api
var once sync.Once

// GetAPIInstance returns a singleton reference
// to the api version.
// It returns the reference to the static api.
func GetAPIInstance() *api {
	once.Do(func() {
		instance = &api{base: 1, sub: 0}
	})
	return instance
}

// SetupResponse provides a convenient way for allowing handlers
// to work with CORS.
// It takes in a ResponseWriter reference to modify the headers with CORS compliant information.
func SetupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}
