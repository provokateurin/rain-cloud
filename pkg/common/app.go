package common

import (
	"fmt"

	"github.com/go-chi/chi/v5"
)

type RegistrationInterface interface {
	GetCapabilities() *string
}

//nolint:revive
func RegisterApp(id string, registration RegistrationInterface, onRegister func(router chi.Router)) {
	fmt.Printf("Registering app: %s\n", id)

	onRegister(router)
}
