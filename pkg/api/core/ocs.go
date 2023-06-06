package coreapi

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/common"
)

//nolint:lll,revive
func (c CoreImpl) OcsGetCapabilities(ctx context.Context, request OcsGetCapabilitiesRequestObject) (OcsGetCapabilitiesResponseObject, error) {
	return OcsGetCapabilities200JSONResponse{Ocs: struct {
		Data struct {
			Capabilities map[string]map[string]interface{} `json:"capabilities"`
			Version      struct {
				Edition         string `json:"edition"`
				ExtendedSupport bool   `json:"extendedSupport"`
				Major           int64  `json:"major"`
				Micro           int64  `json:"micro"`
				Minor           int64  `json:"minor"`
				String          string `json:"string"`
			} `json:"version"`
		} `json:"data"`
		Meta OCSMeta `json:"meta"`
	}(struct {
		Data struct {
			Capabilities map[string]map[string]interface{} `json:"capabilities"`
			Version      struct {
				Edition         string `json:"edition"`
				ExtendedSupport bool   `json:"extendedSupport"`
				Major           int64  `json:"major"`
				Micro           int64  `json:"micro"`
				Minor           int64  `json:"minor"`
				String          string `json:"string"`
			} `json:"version"`
		}
		Meta OCSMeta
	}{Data: struct {
		Capabilities map[string]map[string]interface{} `json:"capabilities"`
		Version      struct {
			Edition         string `json:"edition"`
			ExtendedSupport bool   `json:"extendedSupport"`
			Major           int64  `json:"major"`
			Micro           int64  `json:"micro"`
			Minor           int64  `json:"minor"`
			String          string `json:"string"`
		} `json:"version"`
	}(struct {
		Capabilities map[string]map[string]interface{}
		Version      struct {
			Edition         string `json:"edition"`
			ExtendedSupport bool   `json:"extendedSupport"`
			Major           int64  `json:"major"`
			Micro           int64  `json:"micro"`
			Minor           int64  `json:"minor"`
			String          string `json:"string"`
		}
	}{Capabilities: map[string]map[string]interface{}{}, Version: struct {
		Edition         string `json:"edition"`
		ExtendedSupport bool   `json:"extendedSupport"`
		Major           int64  `json:"major"`
		Micro           int64  `json:"micro"`
		Minor           int64  `json:"minor"`
		String          string `json:"string"`
	}(struct {
		Edition         string
		ExtendedSupport bool
		Major           int64
		Micro           int64
		Minor           int64
		String          string
	}{
		Edition:         common.Edition,
		ExtendedSupport: common.ExtendedSupport,
		Major:           common.VersionMajor,
		Micro:           common.VersionMicro,
		Minor:           common.VersionMinor,
		String:          common.VersionString,
	})}), Meta: common.DummyOCSMeta})}, nil
}
