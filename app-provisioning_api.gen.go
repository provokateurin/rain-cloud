//go:build provisioning_api

package main

import (
	"github.com/go-chi/chi/v5"
	provisioning_apiapi "github.com/provokateurin/rain-cloud/pkg/api/provisioning_api"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func init() {
	common.RegisterApp("provisioning_api", func(router chi.Router) {
		var provisioning_apiAPI provisioning_apiapi.ProvisioningApiImpl
		provisioning_apiapi.HandlerFromMux(provisioning_apiapi.NewStrictHandler(&provisioning_apiAPI, nil), router)
	})
}
