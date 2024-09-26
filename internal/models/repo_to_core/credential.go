package repo_to_core

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func RepoCredentialToCoreCredential(credential repo.Credential) core.Credential {
	return core.Credential{
		Id:         core.CredentialId(credential.CredentialId),
		Credential: core.CredentialVal(credential.Code),
	}
}
