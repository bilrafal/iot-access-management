package repo_to_core

import (
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/repo"
)

func RepoUserToCoreUser(user repo.User) core.User {
	return core.User{
		Id:   core.UserId(user.UserId),
		Name: user.Name,
	}
}
