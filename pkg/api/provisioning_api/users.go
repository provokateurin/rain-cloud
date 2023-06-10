//nolint:revive,stylecheck
package provisioning_api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/provokateurin/rain-cloud/pkg/common"
)

//nolint:lll,revive
func (p ProvisioningApiAPI) UsersGetCurrentUser(ctx context.Context, request UsersGetCurrentUserRequestObject) (UsersGetCurrentUserResponseObject, error) {
	quota, err := json.Marshal(0)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal quota: %w", err)
	}
	scope := "v2-private"

	return UsersGetCurrentUser200JSONResponse{Ocs: struct {
		Data UserDetails `json:"data"`
		Meta OCSMeta     `json:"meta"`
	}(struct {
		Data UserDetails
		Meta OCSMeta
	}{Data: UserDetails{
		AdditionalMail:      nil,
		AdditionalMailScope: &[]string{scope},
		Address:             "",
		AddressScope:        &scope,
		AvatarScope:         &scope,
		Backend:             "",
		BackendCapabilities: struct {
			SetDisplayName bool `json:"setDisplayName"`
			SetPassword    bool `json:"setPassword"`
		}{},
		Biography:           "",
		BiographyScope:      &scope,
		DisplayName:         "Admin",
		Displayname:         "Admin",
		DisplaynameScope:    &scope,
		Email:               nil,
		EmailScope:          &scope,
		Enabled:             nil,
		Fediverse:           common.Pointer(""),
		FediverseScope:      &scope,
		Groups:              nil,
		Headline:            "",
		HeadlineScope:       &scope,
		Id:                  "admin",
		Language:            "",
		LastLogin:           0,
		Locale:              "en_US",
		NotifyEmail:         nil,
		Organisation:        "",
		OrganisationScope:   &scope,
		Phone:               "",
		PhoneScope:          &scope,
		ProfileEnabled:      "",
		ProfileEnabledScope: &scope,
		Quota: struct {
			Free     *int64                  `json:"free"`
			Quota    UserDetails_Quota_Quota `json:"quota"`
			Relative *float32                `json:"relative"`
			Total    *int64                  `json:"total"`
			Used     int64                   `json:"used"`
		}{
			Free: common.Pointer(int64(0)),
			Quota: UserDetails_Quota_Quota{
				quota,
			},
			Relative: common.Pointer(float32(0)),
			Total:    common.Pointer(int64(0)),
		},
		Role:            "",
		RoleScope:       &scope,
		StorageLocation: nil,
		Subadmin:        nil,
		Twitter:         "",
		TwitterScope:    &scope,
		Website:         "",
		WebsiteScope:    &scope,
	}, Meta: common.DummyOCSMeta})}, nil
}
