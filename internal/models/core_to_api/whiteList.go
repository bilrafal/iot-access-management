package core_to_api

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
)

func CoreWhiteListToApiWhiteList(whiteList core.WhiteList) api.WhiteListCreateRequest {
	return *api.NewWhiteListCreateRequest(string(whiteList.DoorId), string(whiteList.CredentialId))
}

func CoreAccessRequestToApiAccessRequest(request core.AccessRequest) api.AccessRequest {
	return *api.NewAccessRequest(string(request.DoorId), string(request.Credential))
}
