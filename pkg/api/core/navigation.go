package core

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/common"
)

//nolint:lll,revive
func (c CoreAPI) NavigationGetAppsNavigation(ctx context.Context, request NavigationGetAppsNavigationRequestObject) (NavigationGetAppsNavigationResponseObject, error) {
	return NavigationGetAppsNavigation200JSONResponse{Ocs: struct {
		Data []NavigationEntry `json:"data"`
		Meta OCSMeta           `json:"meta"`
	}(struct {
		Data []NavigationEntry
		Meta OCSMeta
	}{Data: []NavigationEntry{}, Meta: common.DummyOCSMeta})}, nil
}
