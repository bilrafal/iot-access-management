package core_to_api

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
)

func CoreCredentialToApiCredentialGetResponse(coreCredential core.Credential) api.CredentialResponse {
	return api.CredentialResponse{
		Id:         string(coreCredential.Id),
		Credential: string(coreCredential.Credential),
	}
}

func CoreCredentialToApiCreateCredentialResponse(coreCredential core.Credential) api.CredentialCreateResponse {
	return api.CredentialCreateResponse{
		Id: string(coreCredential.Id),
	}
}
