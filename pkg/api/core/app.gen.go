package coreapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

type CoreAPI struct{}

var _ StrictServerInterface = (*CoreAPI)(nil)

type CoreRegistration struct {}

var _ common.RegistrationInterface = (*CoreRegistration)(nil)

func init() {
    var registration CoreRegistration
	common.RegisterApp("core", registration, func(router chi.Router) {
		var api CoreAPI
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
