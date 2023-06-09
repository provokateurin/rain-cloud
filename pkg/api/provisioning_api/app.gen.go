//go:build app

package provisioning_apiapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func init() {
	common.RegisterApp("provisioning_api", func(router chi.Router) {
		var api ProvisioningApiImpl
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
