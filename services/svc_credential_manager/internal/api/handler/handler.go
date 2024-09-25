package handler

import (
	"encoding/json"
	"fmt"
	"iot-access-management/internal/models/api_to_core"
	"iot-access-management/internal/models/core_to_api"
	"iot-access-management/internal/util/http_helper"
	api_context "iot-access-management/services/svc_credential_manager/internal/api/api_context"
	"iot-access-management/services/svc_credential_manager/internal/core_manager"
	"log/slog"
	"net/http"
)

type CredentialHandler struct {
	credentialsManager core_manager.CredentialManager
}

func NewCredentialHandler(cm core_manager.CredentialManager) *CredentialHandler {
	return &CredentialHandler{
		credentialsManager: cm,
	}
}

func (ch *CredentialHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	userRequest, err := api_context.GetUserCreatePayloadFromBody(w, r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Json not valid, expected Json: [%+v]; error: ", userRequest, err)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := api_to_core.ApiCreateUserToCoreUser(*userRequest)

	newUser, err := ch.credentialsManager.CreateUser(user)
	if err != nil {
		slog.Error("failed to insert:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(newUser)
	if err != nil {
		slog.Error("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ch *CredentialHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	userId := api_context.GetUserIdFromUrlParam(r)

	//TODO: switch depending if NotFound or InternalError
	coreUser, err := ch.credentialsManager.GetUser(userId)
	if err != nil {
		slog.Error("failed to get user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiUser := core_to_api.CoreUserToApiUser(*coreUser)

	http_helper.RespondOkWithBody(w, apiUser)
}
