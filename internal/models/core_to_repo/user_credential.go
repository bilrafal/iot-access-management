package core_to_repo

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func CoreUserCredentialToRepoUserCredential(userCredential core.UserCredential) repo.UserCredential {
	return *repo.NewUserCredential(string(userCredential.UserId), string(userCredential.CredentialId))

}
