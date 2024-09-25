package handler

import (
	"fmt"
	"iot-access-management/internal/models/api_to_core"
	"iot-access-management/internal/models/core_to_api"
	"iot-access-management/internal/util/http_helper"
	"iot-access-management/services/svc_credential_manager/internal/api/api_context"
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

	userRequest, err := api_context.GetUserCreatePayloadFromBody(r)
	if err != nil {
		http_helper.RespondBadRequestWithError(w,
			fmt.Sprintf("json not valid, expected Json: [%+v]; error: [%v]", userRequest, err))
		return
	}

	user := api_to_core.ApiCreateUserToCoreUser(*userRequest)
	newUser, err := ch.credentialsManager.CreateUser(user)

	if err != nil {
		http_helper.RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to insert: [%v]", err))
		return
	}

	http_helper.RespondWithStatusCreatedAndBody(w, newUser)
}

func (ch *CredentialHandler) CreateCredential(w http.ResponseWriter, r *http.Request) {

	credentialRequest, err := api_context.GetCredentialCreatePayloadFromBody(r)
	if err != nil {
		http_helper.RespondBadRequestWithError(w,
			fmt.Sprintf("json not valid, expected Json: [%+v]; error: [%v]", credentialRequest, err))
		return
	}

	credential := api_to_core.ApiCreateCredentialToCoreCredential(*credentialRequest)
	newCredential, err := ch.credentialsManager.CreateCredential(string(credential.Credential))

	if err != nil {
		http_helper.RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to insert: [%v]", err))
		return
	}

	http_helper.RespondWithStatusCreatedAndBody(w, newCredential)
}

func (ch *CredentialHandler) AssignCredentialToUser(w http.ResponseWriter, r *http.Request) {

	credentialRequest, err := api_context.GetAssignCredentialToUserPayloadFromBody(r)
	if err != nil {
		http_helper.RespondBadRequestWithError(w,
			fmt.Sprintf("json not valid, expected Json: [%+v]; error: [%v]", credentialRequest, err))
		return
	}

	coreCredential := api_to_core.ApiAssignCredentialToUserRequestToCoreUserCredential(*credentialRequest)
	err = ch.credentialsManager.AssignCredentialToUser(coreCredential.UserId, coreCredential.CredentialId)

	if err != nil {
		http_helper.RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to insert: [%v]", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
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

func (ch *CredentialHandler) GetUserCredentials(w http.ResponseWriter, r *http.Request) {

	userId := api_context.GetUserIdFromUrlParam(r)

	//TODO: switch depending if NotFound or InternalError
	coreUserCredential, err := ch.credentialsManager.GetUserCredentials(userId)
	if err != nil {
		slog.Error("failed to get user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http_helper.RespondOkWithBody(w, coreUserCredential)
}
