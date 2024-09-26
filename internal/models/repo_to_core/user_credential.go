package repo_to_core

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func RepoUserCredentialToCoreUserCredential(credential repo.UserCredential) core.UserCredential {
	return core.UserCredential{
		Id:           credential.Id,
		UserId:       core.UserId(credential.UserId),
		CredentialId: core.CredentialId(credential.CredentialId),
	}
}
