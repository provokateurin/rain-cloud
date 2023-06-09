package coreapi

import "context"

//nolint:lll,revive
func (c CoreImpl) ClientFlowLoginV2Init(ctx context.Context, request ClientFlowLoginV2InitRequestObject) (ClientFlowLoginV2InitResponseObject, error) {
	return ClientFlowLoginV2Init200JSONResponse{
		Login: "https://example.com",
		Poll: struct {
			Endpoint string `json:"endpoint"`
			Token    string `json:"token"`
		}(struct {
			Endpoint string
			Token    string
		}{Endpoint: "http://localhost/index.php/login/v2/poll", Token: "stub"}),
	}, nil
}

//nolint:lll,revive
func (c CoreImpl) ClientFlowLoginV2Poll(ctx context.Context, request ClientFlowLoginV2PollRequestObject) (ClientFlowLoginV2PollResponseObject, error) {
	return ClientFlowLoginV2Poll200JSONResponse{
		AppPassword: "stub",
		LoginName:   "admin",
		Server:      "http://localhost",
	}, nil
}
