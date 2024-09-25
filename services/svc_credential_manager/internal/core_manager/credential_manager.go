package core_manager

import (
	"iot-access-management/internal/models/core"
)

type CredentialManager interface {
	CreateUser(user core.User) (*core.User, error)
	GetUser(id core.UserId) (*core.User, error)
	CreateCredential(accessCode string) (core.CredentialId, error)
	//AssignCredentialToUser(userId string, credId core.CredentialId) error
	//AuthorizeUserOnDoor(credId core.CredentialId) error
	//RevokeAuthorization(credId core.CredentialId) error
}
