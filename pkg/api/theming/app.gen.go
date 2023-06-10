package themingapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

type ThemingAPI struct{}

var _ StrictServerInterface = (*ThemingAPI)(nil)

type ThemingRegistration struct {}

var _ common.RegistrationInterface = (*ThemingRegistration)(nil)

func init() {
    var registration ThemingRegistration
	common.RegisterApp("theming", registration, func(router chi.Router) {
		var api ThemingAPI
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
