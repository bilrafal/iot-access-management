package api_context

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
	"net/http"
)

const (
	UrlUserId         = "id"
	UrlDoorId         = "door-id"
	UrlCredentialId   = "credential-id"
	UrlCredentialCode = "code"
)

func GetUserIdFromUrlParam(r *http.Request) core.UserId {
	return core.UserId(chi.URLParam(r, UrlUserId))
}
func GetCredentialIdFromUrlParam(r *http.Request) core.CredentialId {
	return core.CredentialId(chi.URLParam(r, UrlCredentialId))
}
func GetCredentialCodeFromUrlParam(r *http.Request) core.CredentialVal {
	return core.CredentialVal(chi.URLParam(r, UrlCredentialCode))
}
func GetDoorIdFromUrlParam(r *http.Request) core.DoorId {
	return core.DoorId(chi.URLParam(r, UrlDoorId))
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

func GetAssignCredentialToUserPayloadFromBody(r *http.Request) (*api.AssignCredentialToUserRequest, error) {
	var assignCredRequest api.AssignCredentialToUserRequest
	if err := json.NewDecoder(r.Body).Decode(&assignCredRequest); err != nil {
		return nil, err
	}

	return &assignCredRequest, nil
}
