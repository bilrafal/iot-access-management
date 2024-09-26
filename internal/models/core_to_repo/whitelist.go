package core_to_repo

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func CoreWhiteListToRepoWhiteList(whiteList core.WhiteList) repo.WhiteList {
	return *repo.NewWhiteList(whiteList.Id, string(whiteList.DoorId), string(whiteList.CredentialId))
}
