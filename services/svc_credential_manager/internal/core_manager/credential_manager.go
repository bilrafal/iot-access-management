package core_manager

import (
	"iot-access-management/internal/error/trace_error"
	"iot-access-management/internal/models/core"
)

var (
	ErrCredentialNotFound = trace_error.NewTraceError("CREDENTIAL_NOT_FOUND")
	ErrUserNotFound       = trace_error.NewTraceError("USER_NOT_FOUND")
	ErrUnexpected         = trace_error.NewTraceError("ERR_UNEXPECTED")
)

type CredentialManager interface {
	CreateUser(user core.User) (*core.User, *trace_error.TraceError)
	GetUser(id core.UserId) (*core.User, *trace_error.TraceError)
	CreateCredential(accessCode string) (core.CredentialId, *trace_error.TraceError)
	AssignCredentialToUser(userId core.UserId, credId core.CredentialId) *trace_error.TraceError
	GetUserCredentials(userId core.UserId) ([]*core.UserCredential, *trace_error.TraceError)
	AuthorizeUserOnDoor(doorId core.DoorId, credId core.CredentialId) *trace_error.TraceError
	RevokeAuthorization(doorId core.DoorId, credId core.CredentialId) *trace_error.TraceError
	GetCredentialIdByCode(credentialCode core.CredentialVal) (*core.Credential, *trace_error.TraceError)
}
