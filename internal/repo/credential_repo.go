package repo

import (
	"iot-access-management/internal/models/repo"
)

type RepoCredential interface {
	AddUser(user repo.User) error
	GetUser(userId string) (*repo.User, error)
	AddCredential(credential repo.Credential) error
	GetCredential(credentialId string) (*repo.Credential, error)
	AddUserCredential(userCredential repo.UserCredential) error
	GetUserCredentials(userId string) ([]*repo.UserCredential, error)
}
