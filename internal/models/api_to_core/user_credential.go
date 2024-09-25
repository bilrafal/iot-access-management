package api_to_core

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
)

func ApiAssignCredentialToUserRequestToCoreUserCredential(apiUserCredential api.AssignCredentialToUserRequest) core.UserCredential {
	return core.UserCredential{
		UserId:       core.UserId(apiUserCredential.UserId),
		CredentialId: core.CredentialId(apiUserCredential.CredentialId),
	}
}
