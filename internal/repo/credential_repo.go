package repo

import (
	"iot-access-management/internal/models/repo"
)

type RepoCredential interface {
	AddUser(user repo.User) error
	GetUser(userId string) (*repo.User, error)
	AddCredential(credential repo.Credential) error
}
