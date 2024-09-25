package in_memory

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"iot-access-management/internal/db"
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

func (dbEng *InMemoryDb) Get(table db.TableName, keys db.KeySet) (interface{}, error) {
	var err error
	var data interface{}
	txn := dbEng.dbEngine.Txn(false)
	defer txn.Abort()

	switch table {
	case db.UserTableName:
		id, ok := keys[db.KeyName(strings.ToLower(db.UserIdFieldName.String()))]
		if !ok {
			return nil, errors.New("invalid keys")
		}

		data, err = txn.First(db.UserTableName.String(), strings.ToLower(db.UserIdFieldName.String()), fmt.Sprintf("%s", id))
		if err != nil {
			return nil, db.ErrDbNotFound
		}
	case db.CredentialTableName:
		id, ok := keys[db.KeyName(strings.ToLower(db.CredentialIdFieldName.String()))]
		if !ok {
			return nil, errors.New("invalid keys")
		}

		data, err = txn.First(db.CredentialTableName.String(), strings.ToLower(db.CredentialIdFieldName.String()), fmt.Sprintf("%s", id))
		if err != nil {
			return nil, db.ErrDbNotFound
		}
	case db.UserCredentialRelTableName:
		userId, ok := keys[db.KeyName(strings.ToLower(db.UserIdFieldName.String()))]
		if !ok {
			return nil, errors.New("invalid keys")
		}

		dataList, err := txn.Get(db.UserCredentialRelTableName.String(), strings.ToLower(db.UserIdFieldName.String()), fmt.Sprintf("%s", userId))
		if err != nil {
			return nil, db.ErrDbNotFound
		}
		var userCredentials []*repo.UserCredential
		for obj := dataList.Next(); obj != nil; obj = dataList.Next() {
			uc := obj.(*repo.UserCredential)
			userCredentials = append(userCredentials, uc)
		}

		return userCredentials, nil
	default:
		return nil, errors.New("unknown table")
	}

	return data, nil

}

func (dbEng *InMemoryDb) Save(table db.TableName, data interface{}) error {

	txn := dbEng.dbEngine.Txn(true)
	err := txn.Insert(string(table), data)
	txn.Commit()

	return err
}

func (dbEng *InMemoryDb) Delete(table db.TableName, keys db.KeySet) error {
	var err error
	txn := dbEng.dbEngine.Txn(true)

	switch table {
	case db.UserTableName:
		id, ok := keys[db.KeyName(db.UserIdFieldName)]
		if !ok {
			return errors.New("invalid keys")
		}
		idStr, ok := id.(string)
		if !ok {
			return errors.New("invalid keys")
		}
		user := repo.User{
			Id: idStr,
		}
		err = txn.Delete(db.UserTableName.String(), user)
		txn.Commit()
	}

	return err
}

func (dbEng *InMemoryDb) Update(table db.TableName, keys db.KeySet, data interface{}) error {
	//TODO implement me
	panic("implement me")
}
