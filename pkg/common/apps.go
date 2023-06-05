package common

import (
	"fmt"

	"github.com/go-chi/chi/v5"
)

//nolint:gochecknoglobals
var apps []string

func RegisterApp(name string, onRegister func(router chi.Router)) {
	for _, app := range apps {
		if app == name {
			panic(fmt.Sprintf("App already registered: %s", name))
		}
	}
	fmt.Printf("Registering app: %s\n", name)
	apps = append(apps, name)
	onRegister(router)
}
