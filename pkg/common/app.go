package common

import (
	"fmt"

	"github.com/go-chi/chi/v5"
)

//nolint:revive
func RegisterApp(id string, onRegister func(router chi.Router)) {
	fmt.Printf("Registering app: %s\n", id)

	onRegister(router)
}
