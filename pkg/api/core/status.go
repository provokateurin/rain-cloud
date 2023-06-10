package coreapi

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/common"
)

//nolint:revive
func (c CoreAPI) GetStatus(ctx context.Context, request GetStatusRequestObject) (GetStatusResponseObject, error) {
	return GetStatus200JSONResponse{
		Edition:         common.Edition,
		ExtendedSupport: common.ExtendedSupport,
		Installed:       true,
		Maintenance:     false,
		NeedsDbUpgrade:  false,
		Productname:     "Nextcloud",
		Version:         common.VersionString,
		Versionstring:   common.VersionString,
	}, nil
}
