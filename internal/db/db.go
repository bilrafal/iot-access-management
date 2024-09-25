package db

import "iot-access-management/internal/error/trace_error"

type DbType string

const (
	InMemory DbType = "in-memory"
)

type TableName string
type FieldName string

const (
	UserTableName     TableName = "user"
	UserIdFieldName   FieldName = "Id"
	UserNameFieldName FieldName = "Name"

	CredentialTableName     TableName = "credential"
	CredentialIdFieldName   FieldName = "Id"
	CredentialCodeFieldName FieldName = "Code"

	UserCredentialRelTableName TableName = "user-credential"
	UserFkIdFieldName          FieldName = "UserId"
	CredentialFkIdFieldName    FieldName = "CredentialId"

	WhiteListedDoorTableName TableName = "whitelist"
	DoorIdFieldName          FieldName = "Id"
)

var (
	ErrDbNotFound      = trace_error.NewTraceError("DB_NOT_FOUND")
	ErrBadData         = trace_error.NewTraceError("DB_BAD_DATA")
	ErrUnexpected      = trace_error.NewTraceError("DB_UNEXPECTED")
	ErrConnectionError = trace_error.NewTraceError("DB_CONNECTION_FAIL")
)

func (tn TableName) String() string {
	return string(tn)
}
func (fn FieldName) String() string {
	return string(fn)
}

type KeyName string
type KeyValue interface{}
type KeySet map[KeyName]KeyValue

type DbClient interface {
	Get(table TableName, keys KeySet) (interface{}, error)
	Save(table TableName, data interface{}) error
	Delete(table TableName, keys KeySet) error
	Update(table TableName, keys KeySet, data interface{}) error
}
