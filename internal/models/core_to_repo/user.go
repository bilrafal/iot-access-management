package core_to_repo

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func CoreUserToRepoUser(user core.User) repo.User {
	return repo.User{
		Id:   string(user.Id),
		Name: user.Name,
	}
}
