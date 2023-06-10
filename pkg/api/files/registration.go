package files

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/registration"
)

//nolint:lll,revive
func (f FilesRegistration) GetRegistration(ctx context.Context, request registration.GetRegistrationRequestObject) (registration.GetRegistrationResponseObject, error) {
	r, err := registration.GenerateRegistrationResponse("files", nil, &registration.NavigationEntry{
		Id:    "files",
		Name:  "Files",
		Order: 0,
	})
	if err != nil {
		panic(err)
	}

	return r, nil
}
