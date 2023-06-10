package provisioning_apiapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

type ProvisioningApiAPI struct{}

var _ StrictServerInterface = (*ProvisioningApiAPI)(nil)

type ProvisioningApiRegistration struct {}

var _ common.RegistrationInterface = (*ProvisioningApiRegistration)(nil)

func init() {
    var registration ProvisioningApiRegistration
	common.RegisterApp("provisioning_api", registration, func(router chi.Router) {
		var api ProvisioningApiAPI
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
