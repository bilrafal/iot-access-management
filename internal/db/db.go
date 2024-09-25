package db

type DbType string

const (
	InMemory DbType = "in-memory"
)

type TableName string
type FieldName string

const (
	UserTableName TableName = "user"

	UserIdFieldName   FieldName = "Id"
	UserNameFieldName FieldName = "Name"
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
	Get(table TableName, keys KeySet, respData interface{}) (interface{}, error)
	Save(table TableName, data interface{}) error
	Delete(table TableName, keys KeySet) error
	Update(table TableName, keys KeySet, data interface{}) error
}
