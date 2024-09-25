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
		},
	}
}
