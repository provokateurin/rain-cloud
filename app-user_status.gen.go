//go:build user_status

package main

import (
	"github.com/go-chi/chi/v5"
	user_statusapi "github.com/provokateurin/rain-cloud/pkg/api/user_status"
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func init() {
	common.RegisterApp("user_status", func(router chi.Router) {
		var user_statusAPI user_statusapi.UserStatusImpl
		user_statusapi.HandlerFromMux(user_statusapi.NewStrictHandler(&user_statusAPI, nil), router)
	})
}
