package core

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"
)

type CoreAPI struct{}

var _ StrictServerInterface = (*CoreAPI)(nil)

type CoreRegistration struct {}

var _ registration.StrictServerInterface = (*CoreRegistration)(nil)

func init() {
    HandlerFromMux(NewStrictHandler(&CoreAPI{}, nil), common.Router)
    registration.HandlerFromMux(registration.NewStrictHandler(&CoreRegistration{}, nil), common.Router)
}
