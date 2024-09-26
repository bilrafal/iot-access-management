package repo_to_core

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func RepoWhiteListToCoreWhiteList(whiteList repo.WhiteList) core.WhiteList {
	return *core.NewWhiteList(whiteList.Id, core.DoorId(whiteList.DoorId), core.CredentialId(whiteList.CredentialId))
}
