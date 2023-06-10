package core

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/registration"
)

//nolint:lll,revive
func (c CoreRegistration) GetRegistration(ctx context.Context, request registration.GetRegistrationRequestObject) (registration.GetRegistrationResponseObject, error) {
	return registration.GetRegistration200JSONResponse{
		Id: "core",
	}, nil
}
