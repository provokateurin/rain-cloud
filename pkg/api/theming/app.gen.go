//go:build app

package themingapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func init() {
	common.RegisterApp("theming", func(router chi.Router) {
		var api ThemingImpl
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
