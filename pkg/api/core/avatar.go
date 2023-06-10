package coreapi

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/provokateurin/rain-cloud/pkg/service"
)

//nolint:lll,revive
func (c CoreAPI) GuestAvatarGetAvatar(ctx context.Context, request GuestAvatarGetAvatarRequestObject) (GuestAvatarGetAvatarResponseObject, error) {
	size, err := parseSize(request.Size)
	if err != nil {
		return nil, err
	}

	avatar, contentType, err := service.CoreService.GenerateAvatar(request.GuestName, size, false)
	if err != nil {
		return GuestAvatarGetAvatar500Response{}, fmt.Errorf("failed to generate avatar: %w", err)
	}

	return GuestAvatarGetAvatar200AsteriskResponse{
		Body:          bytes.NewReader(avatar),
		ContentLength: int64(len(avatar)),
		ContentType:   contentType,
	}, nil
}

//nolint:lll,revive
func (c CoreAPI) GuestAvatarGetAvatarDark(ctx context.Context, request GuestAvatarGetAvatarDarkRequestObject) (GuestAvatarGetAvatarDarkResponseObject, error) {
	size, err := parseSize(request.Size)
	if err != nil {
		return nil, err
	}

	avatar, contentType, err := service.CoreService.GenerateAvatar(request.GuestName, size, true)
	if err != nil {
		return GuestAvatarGetAvatarDark500Response{}, fmt.Errorf("failed to generate avatar: %w", err)
	}

	return GuestAvatarGetAvatarDark200AsteriskResponse{
		Body:          bytes.NewReader(avatar),
		ContentLength: int64(len(avatar)),
		ContentType:   contentType,
	}, nil
}

//nolint:lll,revive
func (c CoreAPI) AvatarGetAvatar(ctx context.Context, request AvatarGetAvatarRequestObject) (AvatarGetAvatarResponseObject, error) {
	avatar, contentType, err := service.CoreService.GenerateAvatar(request.UserId, int(request.Size), false)
	if err != nil {
		return AvatarGetAvatar404JSONResponse{}, fmt.Errorf("failed to generate avatar: %w", err)
	}

	return AvatarGetAvatar200AsteriskResponse{
		Body: bytes.NewReader(avatar),
		Headers: AvatarGetAvatar200ResponseHeaders{
			XNCIsCustomAvatar: 0,
		},
		ContentLength: int64(len(avatar)),
		ContentType:   contentType,
	}, nil
}

//nolint:lll,revive
func (c CoreAPI) AvatarGetAvatarDark(ctx context.Context, request AvatarGetAvatarDarkRequestObject) (AvatarGetAvatarDarkResponseObject, error) {
	avatar, contentType, err := service.CoreService.GenerateAvatar(request.UserId, int(request.Size), true)
	if err != nil {
		return AvatarGetAvatarDark404JSONResponse{}, fmt.Errorf("failed to generate avatar: %w", err)
	}

	return AvatarGetAvatarDark200AsteriskResponse{
		Body: bytes.NewReader(avatar),
		Headers: AvatarGetAvatarDark200ResponseHeaders{
			XNCIsCustomAvatar: 0,
		},
		ContentLength: int64(len(avatar)),
		ContentType:   contentType,
	}, nil
}

func parseSize(size string) (int, error) {
	s, err := strconv.Atoi(size)
	if err != nil {
		return 0, fmt.Errorf("failed to parse size: %w", err)
	}

	return s, nil
}
