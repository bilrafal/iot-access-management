package api_to_core

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
)

func ApiCreateUserToCoreUser(apiUser api.UserCreateRequest) core.User {
	return core.User{
		Name: apiUser.Name,
	}
}
