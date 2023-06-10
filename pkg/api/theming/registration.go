package theming

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/registration"
)

//nolint:lll,revive
func (c ThemingRegistration) GetRegistration(ctx context.Context, request registration.GetRegistrationRequestObject) (registration.GetRegistrationResponseObject, error) {
	return registration.GetRegistration200JSONResponse{
		Id: "theming",
	}, nil
}
