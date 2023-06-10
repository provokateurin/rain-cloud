//nolint:revive,stylecheck
package user_status

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/registration"
)

//nolint:lll,revive
func (c UserStatusRegistration) GetRegistration(ctx context.Context, request registration.GetRegistrationRequestObject) (registration.GetRegistrationResponseObject, error) {
	return registration.GetRegistration200JSONResponse{
		Id: "user_status",
	}, nil
}
