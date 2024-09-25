package repo

import (
	"iot-access-management/internal/models/repo"
)

type RepoCredential interface {
	AddUser(user repo.User) error
	GetUser(userId repo.UserId) (*repo.User, error)
	//AddCredential(accessCode string) error
}
