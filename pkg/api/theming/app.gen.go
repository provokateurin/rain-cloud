package theming

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"

	"github.com/go-chi/chi/v5"
)

type ThemingAPI struct{}

var _ StrictServerInterface = (*ThemingAPI)(nil)

type ThemingRegistration struct {}

var _ registration.StrictServerInterface = (*ThemingRegistration)(nil)

func init() {
	common.RegisterApp("theming", func(router chi.Router) {
		HandlerFromMux(NewStrictHandler(&ThemingAPI{}, nil), router)
        registration.HandlerFromMux(registration.NewStrictHandler(&ThemingRegistration{}, nil), router)
	})
}
