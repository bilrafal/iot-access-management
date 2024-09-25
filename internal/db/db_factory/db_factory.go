package db_factory

import (
	"context"
	"errors"
	"fmt"
	"iot-access-management/internal/db"
	in_memory "iot-access-management/internal/db/in-memory"
)

var ErrDbTypeNonValid = errors.New(fmt.Sprintf("Non valid db type. Valid values are [%q]", db.InMemory))

type DbFactory interface {
	GetDbClient() (db.DbClient, error)
}

type DbFactorySimple struct {
	ctx    context.Context
	dbType db.DbType
}

func NewDbFactorySimple(
	ctx context.Context, dbType db.DbType,
) DbFactory {
	return &DbFactorySimple{
		ctx:    ctx,
		dbType: dbType,
	}
}

func (dbf DbFactorySimple) GetDbClient() (db.DbClient, error) {

	switch dbf.dbType {
	case db.InMemory:
		return in_memory.NewInMemoryDb(), nil
	default:
		return nil, ErrDbTypeNonValid
	}
}
