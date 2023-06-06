package coreapi

import (
	"context"
)

//nolint:lll,revive,stylecheck
func (c CoreImpl) NavigationGetAppsNavigation(ctx context.Context, request NavigationGetAppsNavigationRequestObject) (NavigationGetAppsNavigationResponseObject, error) {
	return NavigationGetAppsNavigation200JSONResponse{Ocs: struct {
		Data []NavigationEntry `json:"data"`
		Meta OCSMeta           `json:"meta"`
	}(struct {
		Data []NavigationEntry
		Meta OCSMeta
	}{Data: []NavigationEntry{}, Meta: dummyOCSMeta})}, nil
}

//nolint:lll,revive,stylecheck
func (c CoreImpl) NavigationGetSettingsNavigation(ctx context.Context, request NavigationGetSettingsNavigationRequestObject) (NavigationGetSettingsNavigationResponseObject, error) {
	panic("implement me")
}
