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

func GetUserCreatePayloadFromBody(w http.ResponseWriter, r *http.Request) (*api.UserCreateRequest, error) {
	var userRequest api.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		return nil, err
	}

	return &userRequest, nil
}
