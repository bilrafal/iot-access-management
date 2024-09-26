package repo

import (
	"iot-access-management/internal/error/trace_error"
	"iot-access-management/internal/models/repo"
)

type RepoCredential interface {
	AddUser(user repo.User) *trace_error.TraceError
	GetUser(userId string) (*repo.User, *trace_error.TraceError)
	AddCredential(credential repo.Credential) *trace_error.TraceError
	GetCredential(credentialId string) (*repo.Credential, *trace_error.TraceError)
	AddUserCredential(userCredential repo.UserCredential) *trace_error.TraceError
	GetUserCredentials(userId string) ([]*repo.UserCredential, *trace_error.TraceError)
	ListCredentials() ([]repo.Credential, *trace_error.TraceError)
}
