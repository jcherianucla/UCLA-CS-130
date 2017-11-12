package config

import (
	"fmt"
	"sync"
)

type api struct {
	base int
	sub  int
}

func (api *api) UpVersion() {
	if api.sub < 9 {
		api.sub += 1
	} else {
		api.base += 1
		api.sub = 0
	}
}

func (api *api) Gen(uri string) string {
	return fmt.Sprintf("/api/v%d.%d%s", api.base, api.sub, uri)
}

var instance *api
var once sync.Once

func GetAPIInstance() *api {
	once.Do(func() {
		instance = &api{base: 1, sub: 0}
	})
	return instance
}
