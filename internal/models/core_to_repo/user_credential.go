package core_to_repo

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func CoreUserCredentialToRepoUserCredential(userCredential core.UserCredential) repo.UserCredential {
	return *repo.NewUserCredential(userCredential.Id, string(userCredential.UserId), string(userCredential.CredentialId))
}
