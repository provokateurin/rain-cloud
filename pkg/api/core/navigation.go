package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/service"
)

//nolint:lll,revive
func (c CoreAPI) NavigationGetAppsNavigation(ctx context.Context, request NavigationGetAppsNavigationRequestObject) (NavigationGetAppsNavigationResponseObject, error) {
	navigations, err := service.CoreService.GetNavigations(ctx)
	if err != nil {
		panic(err)
	}

	//nolint:prealloc
	var entries []NavigationEntry
	for _, navigation := range navigations {
		orderBytes, err := json.Marshal(navigation.Order)
		if err != nil {
			panic(err)
		}
		entries = append(entries, NavigationEntry{
			Active:  false,
			Classes: "",
			Href:    fmt.Sprintf("/apps/%s", navigation.Id),
			Icon:    "",
			Id:      navigation.Id,
			Name:    navigation.Name,
			Order: NavigationEntry_Order{
				orderBytes,
			},
			Type:   "link",
			Unread: 0,
		})
	}

	return NavigationGetAppsNavigation200JSONResponse{Ocs: struct {
		Data []NavigationEntry `json:"data"`
		Meta OCSMeta           `json:"meta"`
	}(struct {
		Data []NavigationEntry
		Meta OCSMeta
	}{Data: entries, Meta: common.DummyOCSMeta})}, nil
}
