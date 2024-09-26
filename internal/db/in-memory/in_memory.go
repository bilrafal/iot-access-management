package in_memory

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"iot-access-management/internal/db"
	"iot-access-management/internal/error/trace_error"
	"iot-access-management/internal/models/repo"
	"strings"
)

type InMemoryDb struct {
	dbEngine *memdb.MemDB
}

func NewInMemoryDb() *InMemoryDb {
	schema := createSchema()
	dbEngine, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	return &InMemoryDb{dbEngine: dbEngine}
}

func (dbEng *InMemoryDb) Get(table db.TableName, keys db.KeySet) (interface{}, *trace_error.TraceError) {
	var err error
	var data interface{}
	txn := dbEng.dbEngine.Txn(false)
	defer txn.Abort()

	switch table {
	case db.UserTableName:
		id, ok := keys[db.KeyName(strings.ToLower(db.UserFkIdFieldName.String()))]
		if !ok {
			return nil, db.ErrUnexpected.From(errors.New("invalid keys"))
		}

		data, err = txn.First(db.UserTableName.String(), strings.ToLower(db.IdFieldName.String()), fmt.Sprintf("%s", id))
		if err != nil {
			return nil, db.ErrDbNotFound.From(err)
		}
	case db.CredentialTableName:
		id, ok := keys[db.KeyName(strings.ToLower(db.IdFieldName.String()))]
		if !ok {
			return nil, db.ErrUnexpected.From(errors.New("invalid keys"))
		}

		data, err = txn.First(db.CredentialTableName.String(), strings.ToLower(db.IdFieldName.String()), fmt.Sprintf("%s", id))
		if err != nil {
			return nil, db.ErrDbNotFound.From(err)
		}
	case db.UserCredentialRelTableName:
		//List all
		dataList, err := txn.Get(db.UserCredentialRelTableName.String(), strings.ToLower(db.IdFieldName.String()))
		if err != nil {
			return nil, db.ErrDbNotFound.Generate()
		}
		var userCredentials []*repo.UserCredential

		for obj := dataList.Next(); obj != nil; obj = dataList.Next() {
			uc := obj.(*repo.UserCredential)
			userCredentials = append(userCredentials, uc)
		}

		return userCredentials, nil

	case db.WhiteListedDoorTableName:
		//List all
		dataList, err := txn.Get(db.WhiteListedDoorTableName.String(), strings.ToLower(db.IdFieldName.String()))
		if err != nil {
			return nil, db.ErrDbNotFound
		}

		var userCredentials []interface{}

		for obj := dataList.Next(); obj != nil; obj = dataList.Next() {
			userCredentials = append(userCredentials, obj)
		}

		if len(userCredentials) == 0 {
			return nil, db.ErrDbNotFound
		}
		return userCredentials, nil
	default:
		return nil, db.ErrUnexpected.From(errors.New("unknown table"))
	}

	return data, nil

}

func (dbEng *InMemoryDb) Save(table db.TableName, data interface{}) *trace_error.TraceError {

	txn := dbEng.dbEngine.Txn(true)
	err := txn.Insert(string(table), data)
	txn.Commit()

	if err != nil {
		return db.ErrUnexpected.From(err)
	}
	return nil
}

func (dbEng *InMemoryDb) Delete(table db.TableName, item interface{}) *trace_error.TraceError {
	var err error
	txn := dbEng.dbEngine.Txn(true)

	switch table {
	case db.WhiteListedDoorTableName:
		err = txn.Delete(db.WhiteListedDoorTableName.String(), item)
		txn.Commit()
	default:
		return db.ErrUnexpected.From(errors.New("unknown table"))
	}

	if err != nil {
		return db.ErrUnexpected.From(err)
	}
	return nil
}

func (dbEng *InMemoryDb) Update(table db.TableName, keys db.KeySet, data interface{}) *trace_error.TraceError {
	//TODO implement me
	panic("implement me")
}

func (dbEng *InMemoryDb) List(table db.TableName) (interface{}, *trace_error.TraceError) {
	txn := dbEng.dbEngine.Txn(false)
	defer txn.Abort()

	dataList, err := txn.Get(table.String(), strings.ToLower(db.IdFieldName.String()))
	if err != nil {
		return nil, db.ErrDbNotFound
	}

	var result []interface{}

	for obj := dataList.Next(); obj != nil; obj = dataList.Next() {
		result = append(result, obj)
	}

	if len(result) == 0 {
		return nil, db.ErrDbNotFound
	}
	return result, nil
}
