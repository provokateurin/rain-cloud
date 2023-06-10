package provisioning_api

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"

	"github.com/go-chi/chi/v5"
)

type ProvisioningApiAPI struct{}

var _ StrictServerInterface = (*ProvisioningApiAPI)(nil)

type ProvisioningApiRegistration struct {}

var _ registration.StrictServerInterface = (*ProvisioningApiRegistration)(nil)

func init() {
	common.RegisterApp("provisioning_api", func(router chi.Router) {
		HandlerFromMux(NewStrictHandler(&ProvisioningApiAPI{}, nil), router)
        registration.HandlerFromMux(registration.NewStrictHandler(&ProvisioningApiRegistration{}, nil), router)
	})
}
