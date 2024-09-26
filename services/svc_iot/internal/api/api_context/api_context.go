package api_context

import (
	"encoding/json"
	"iot-access-management/internal/models/api"
	"net/http"
)

func GetWhiteListCreatePayloadFromBody(r *http.Request) (*api.WhiteListCreateRequest, error) {
	var whiteListRequest api.WhiteListCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&whiteListRequest); err != nil {
		return nil, err
	}

	return &whiteListRequest, nil
}
func GetAccessRequestPayloadFromBody(r *http.Request) (*api.AccessRequest, error) {
	var accessRequest api.AccessRequest
	if err := json.NewDecoder(r.Body).Decode(&accessRequest); err != nil {
		return nil, err
	}

	return &accessRequest, nil
}
