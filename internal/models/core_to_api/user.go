package core_to_api

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
)

func CoreUserToApiUser(coreUser core.User) api.UserResponse {
	return api.UserResponse{
		Id:   string(coreUser.Id),
		Name: coreUser.Name,
	}
}
