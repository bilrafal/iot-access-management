package core_to_api

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
)

func CoreUserCredentialToApiUserCredential(coreUserCredential core.UserCredential) api.AssignCredentialToUserRequest {
	return api.AssignCredentialToUserRequest{
		UserId:       string(coreUserCredential.UserId),
		CredentialId: string(coreUserCredential.CredentialId),
	}
}
