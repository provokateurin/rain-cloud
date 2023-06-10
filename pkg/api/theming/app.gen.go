package theming

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"
)

type ThemingAPI struct{}

var _ StrictServerInterface = (*ThemingAPI)(nil)

type ThemingRegistration struct {}

var _ registration.StrictServerInterface = (*ThemingRegistration)(nil)

func init() {
    HandlerFromMux(NewStrictHandler(&ThemingAPI{}, nil), common.Router)
    registration.HandlerFromMux(registration.NewStrictHandler(&ThemingRegistration{}, nil), common.Router)
}
