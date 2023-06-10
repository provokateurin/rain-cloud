package user_status

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"

	"github.com/go-chi/chi/v5"
)

type UserStatusAPI struct{}

var _ StrictServerInterface = (*UserStatusAPI)(nil)

type UserStatusRegistration struct {}

var _ registration.StrictServerInterface = (*UserStatusRegistration)(nil)

func init() {
	common.RegisterApp("user_status", func(router chi.Router) {
		HandlerFromMux(NewStrictHandler(&UserStatusAPI{}, nil), router)
        registration.HandlerFromMux(registration.NewStrictHandler(&UserStatusRegistration{}, nil), router)
	})
}
