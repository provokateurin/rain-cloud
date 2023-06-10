//nolint:revive
package provisioning_api

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/registration"
)

//nolint:lll,revive
func (c ProvisioningApiRegistration) GetRegistration(ctx context.Context, request registration.GetRegistrationRequestObject) (registration.GetRegistrationResponseObject, error) {
	return registration.GetRegistration200JSONResponse{
		Id: "provisioning_api",
	}, nil
}
