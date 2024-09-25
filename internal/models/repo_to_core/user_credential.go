package repo_to_core

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func RepoUserCredentialToCoreUserCredential(credential repo.UserCredential) core.UserCredential {
	return core.UserCredential{
		UserId:       core.UserId(credential.Id),
		CredentialId: core.CredentialId(credential.CredentialId),
	}
}
