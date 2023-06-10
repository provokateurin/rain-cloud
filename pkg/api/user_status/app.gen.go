package user_status

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"
)

type UserStatusAPI struct{}

var _ StrictServerInterface = (*UserStatusAPI)(nil)

type UserStatusRegistration struct {}

var _ registration.StrictServerInterface = (*UserStatusRegistration)(nil)

func init() {
    HandlerFromMux(NewStrictHandler(&UserStatusAPI{}, nil), common.Router)
    registration.HandlerFromMux(registration.NewStrictHandler(&UserStatusRegistration{}, nil), common.Router)
}
