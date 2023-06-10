package user_statusapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

type UserStatusAPI struct{}

var _ StrictServerInterface = (*UserStatusAPI)(nil)

type UserStatusRegistration struct {}

var _ common.RegistrationInterface = (*UserStatusRegistration)(nil)

func init() {
    var registration UserStatusRegistration
	common.RegisterApp("user_status", registration, func(router chi.Router) {
		var api UserStatusAPI
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
