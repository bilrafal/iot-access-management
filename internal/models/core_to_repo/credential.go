package core_to_repo

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func CoreCredentialToRepoCredential(credential core.Credential) repo.Credential {
	return *repo.NewCredential(string(credential.Id), string(credential.Credential))

}
