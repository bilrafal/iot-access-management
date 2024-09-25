package in_memory

import (
	"github.com/hashicorp/go-memdb"
	"iot-access-management/internal/db"
	"strings"
	//"strings"
)

func createSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			db.UserTableName.String(): &memdb.TableSchema{
				Name: db.UserTableName.String(),
				Indexes: map[string]*memdb.IndexSchema{
					strings.ToLower(db.UserIdFieldName.String()): &memdb.IndexSchema{
						Name:    strings.ToLower(db.UserIdFieldName.String()),
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: db.UserIdFieldName.String()},
					},
				},
			},
			db.CredentialTableName.String(): &memdb.TableSchema{
				Name: db.CredentialTableName.String(),
				Indexes: map[string]*memdb.IndexSchema{
					strings.ToLower(db.CredentialIdFieldName.String()): &memdb.IndexSchema{
						Name:    strings.ToLower(db.CredentialIdFieldName.String()),
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: db.CredentialIdFieldName.String()},
					},
				},
			},
			db.UserCredentialRelTableName.String(): &memdb.TableSchema{
				Name: db.UserCredentialRelTableName.String(),
				Indexes: map[string]*memdb.IndexSchema{
					strings.ToLower(db.UserIdFieldName.String()): &memdb.IndexSchema{
						Name:    strings.ToLower(db.UserIdFieldName.String()),
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: db.UserIdFieldName.String()},
					},
					strings.ToLower(db.CredentialFkIdFieldName.String()): &memdb.IndexSchema{
						Name:    strings.ToLower(db.CredentialFkIdFieldName.String()),
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: db.CredentialFkIdFieldName.String()},
					},
				},
			},

			db.WhiteListedDoorTableName.String(): &memdb.TableSchema{
				Name: db.WhiteListedDoorTableName.String(),
				Indexes: map[string]*memdb.IndexSchema{
					strings.ToLower(db.UserFkIdFieldName.String()): &memdb.IndexSchema{
						Name:    strings.ToLower(db.UserFkIdFieldName.String()),
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: db.UserFkIdFieldName.String()},
					},
					strings.ToLower(db.DoorIdFieldName.String()): &memdb.IndexSchema{
						Name:    strings.ToLower(db.DoorIdFieldName.String()),
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: db.DoorIdFieldName.String()},
					},
				},
			},
		},
	}
}
