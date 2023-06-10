package registration

import (
	"encoding/json"
	"fmt"
)

//nolint:lll
func GenerateRegistrationResponse(id string, capabilities any, navigation *NavigationEntry) (*GetRegistration200JSONResponse, error) {
	var capabilitiesPointer *map[string]map[string]interface{}
	if capabilities != nil {
		bytes, err := json.Marshal(&capabilities)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal capabilities: %w", err)
		}

		var mapCapabilities map[string]map[string]interface{}
		err = json.Unmarshal(bytes, &mapCapabilities)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal capabilities: %w", err)
		}

		capabilitiesPointer = &mapCapabilities
	}

	return &GetRegistration200JSONResponse{
		Id:           id,
		Capabilities: capabilitiesPointer,
		Navigation:   navigation,
	}, nil
}
