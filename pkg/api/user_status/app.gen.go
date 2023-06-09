//go:build app

package user_statusapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func init() {
	common.RegisterApp("user_status", func(router chi.Router) {
		var api UserStatusImpl
		HandlerFromMux(NewStrictHandler(&api, nil), router)
	})
}
