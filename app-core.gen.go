//go:build core

package main

import (
	"github.com/go-chi/chi/v5"
	coreapi "github.com/provokateurin/rain-cloud/pkg/api/core"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

//nolint:gochecknoinits
func init() {
	common.RegisterApp("core", func(router chi.Router) {
		var coreAPI coreapi.CoreImpl
		coreapi.HandlerFromMux(coreapi.NewStrictHandler(&coreAPI, nil), router)
	})
}
