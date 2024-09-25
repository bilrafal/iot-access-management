package core_manager

import (
	"iot-access-management/internal/error/trace_error"
	"iot-access-management/internal/models/core"
)

var (
	ErrUserNotFound       = trace_error.NewTraceError("USER_NOT_FOUND")
	ErrCredentialNotFound = trace_error.NewTraceError("CREDENTIAL_NOT_FOUND")
)

type CredentialManager interface {
	CreateUser(user core.User) (*core.User, error)
	GetUser(id core.UserId) (*core.User, error)
	CreateCredential(accessCode string) (core.CredentialId, error)
	AssignCredentialToUser(userId core.UserId, credId core.CredentialId) error
	GetUserCredentials(userId core.UserId) ([]*core.UserCredential, error)
	//AuthorizeUserOnDoor(credId core.CredentialId) error
	//RevokeAuthorization(credId core.CredentialId) error
}
