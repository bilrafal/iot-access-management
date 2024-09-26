package api_to_core

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/util"
)

func ApiWhiteListCreateRequestToCoreWhiteList(whiteList api.WhiteListCreateRequest) core.WhiteList {
	return *core.NewWhiteList(util.VoidString, core.DoorId(whiteList.DoorId), core.CredentialId(whiteList.CredentialId))
}

func ApiAccessRequestToCoreAccessRequest(request api.AccessRequest) core.AccessRequest {
	return *core.NewAccessRequest(core.DoorId(request.DoorId), core.CredentialVal(request.Credential))
}
