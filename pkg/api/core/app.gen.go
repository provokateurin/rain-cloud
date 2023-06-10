package core

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"

	"github.com/go-chi/chi/v5"
)

type CoreAPI struct{}

var _ StrictServerInterface = (*CoreAPI)(nil)

type CoreRegistration struct {}

var _ registration.StrictServerInterface = (*CoreRegistration)(nil)

func init() {
	common.RegisterApp("core", func(router chi.Router) {
		HandlerFromMux(NewStrictHandler(&CoreAPI{}, nil), router)
        registration.HandlerFromMux(registration.NewStrictHandler(&CoreRegistration{}, nil), router)
	})
}
