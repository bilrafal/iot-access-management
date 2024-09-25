package api_context

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
	"net/http"
)

const (
	UrlUserId = "id"
)

func GetUserIdFromUrlParam(r *http.Request) core.UserId {
	return core.UserId(chi.URLParam(r, UrlUserId))
}

func GetUserCreatePayloadFromBody(r *http.Request) (*api.UserCreateRequest, error) {
	var userRequest api.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		return nil, err
	}

	return &userRequest, nil
}

func GetCredentialCreatePayloadFromBody(r *http.Request) (*api.CredentialCreateRequest, error) {
	var userRequest api.CredentialCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		return nil, err
	}

	return &userRequest, nil
}
