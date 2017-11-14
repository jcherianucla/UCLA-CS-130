// The utilities package provides additional functionality
// available to all packages within the application.
package utilities

import (
	"fmt"
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
