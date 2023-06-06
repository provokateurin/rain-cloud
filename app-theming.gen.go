//go:build theming

package main

import (
	"github.com/go-chi/chi/v5"
	themingapi "github.com/provokateurin/rain-cloud/pkg/api/theming"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func init() {
	common.RegisterApp("theming", func(router chi.Router) {
		var themingAPI themingapi.ThemingImpl
		themingapi.HandlerFromMux(themingapi.NewStrictHandler(&themingAPI, nil), router)
	})
}
