package db

import "iot-access-management/internal/error/trace_error"

type DbType string

const (
	InMemory DbType = "in-memory"
)

type TableName string
type FieldName string

const (
	IdFieldName             FieldName = "Id"
	UserFkIdFieldName       FieldName = "UserId"
	CredentialFkIdFieldName FieldName = "CredentialId"

	UserTableName     TableName = "user"
	UserNameFieldName FieldName = "Name"

	CredentialTableName     TableName = "credential"
	CredentialCodeFieldName FieldName = "Code"

	UserCredentialRelTableName TableName = "user-credential"

	WhiteListedDoorTableName TableName = "whitelist"
	DoorIdFieldName          FieldName = "DoorId"
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
	Get(table TableName, keys KeySet) (interface{}, *trace_error.TraceError)
	List(table TableName) (interface{}, *trace_error.TraceError)
	Save(table TableName, data interface{}) *trace_error.TraceError
	Delete(table TableName, item interface{}) *trace_error.TraceError
	Update(table TableName, keys KeySet, data interface{}) *trace_error.TraceError
}
