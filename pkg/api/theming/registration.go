package theming

import (
	"context"

	"github.com/provokateurin/rain-cloud/pkg/registration"
)

//nolint:lll,revive
func (c ThemingRegistration) GetRegistration(ctx context.Context, request registration.GetRegistrationRequestObject) (registration.GetRegistrationResponseObject, error) {
	capabilities := PublicCapabilities{
		Theming: struct {
			Background         string `json:"background"`
			BackgroundDefault  bool   `json:"background-default"`
			BackgroundPlain    bool   `json:"background-plain"`
			Color              string `json:"color"`
			ColorElement       string `json:"color-element"`
			ColorElementBright string `json:"color-element-bright"`
			ColorElementDark   string `json:"color-element-dark"`
			ColorText          string `json:"color-text"`
			Favicon            string `json:"favicon"`
			Logo               string `json:"logo"`
			Logoheader         string `json:"logoheader"`
			Name               string `json:"name"`
			Slogan             string `json:"slogan"`
			Url                string `json:"url"`
		}(struct {
			Background         string
			BackgroundDefault  bool
			BackgroundPlain    bool
			Color              string
			ColorElement       string
			ColorElementBright string
			ColorElementDark   string
			ColorText          string
			Favicon            string
			Logo               string
			Logoheader         string
			Name               string
			Slogan             string
			Url                string
		}{
			Background:         "#0082c9",
			BackgroundDefault:  true,
			BackgroundPlain:    true,
			Color:              "#0082c9",
			ColorElement:       "#0082c9",
			ColorElementBright: "#0082c9",
			ColorElementDark:   "#0082c9",
			ColorText:          "#ffffff",
			Favicon:            "",
			Logo:               "",
			Logoheader:         "",
			Name:               "RainCloud",
			Slogan:             "Let the rain fall",
			Url:                "http://localhost",
		}),
	}

	r, err := registration.GenerateRegistrationResponse("theming", capabilities, nil)
	if err != nil {
		panic(err)
	}

	return r, nil
}
