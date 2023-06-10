package provisioning_api

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"
)

type ProvisioningApiAPI struct{}

var _ StrictServerInterface = (*ProvisioningApiAPI)(nil)

type ProvisioningApiRegistration struct {}

var _ registration.StrictServerInterface = (*ProvisioningApiRegistration)(nil)

func init() {
    HandlerFromMux(NewStrictHandler(&ProvisioningApiAPI{}, nil), common.Router)
    registration.HandlerFromMux(registration.NewStrictHandler(&ProvisioningApiRegistration{}, nil), common.Router)
}
