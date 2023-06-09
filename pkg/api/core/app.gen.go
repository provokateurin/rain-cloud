//go:build app

package coreapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func init() {
	common.RegisterApp("core", func(router chi.Router) {
		var api CoreImpl
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
