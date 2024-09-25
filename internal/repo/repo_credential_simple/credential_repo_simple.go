package repo_credential_simple

import (
	"context"
	"errors"
	"iot-access-management/internal/db"
	db_factory "iot-access-management/internal/db/db_factory"
	repo_model "iot-access-management/internal/models/repo"
	"iot-access-management/internal/repo"
	"strings"
)

type RepoCredentialSimple struct {
	ctx      context.Context
	dbEngine db.DbClient
}

func NewRepoCredentialSimple(ctx context.Context, dbType db.DbType) repo.RepoCredential {
	dbFactory := db_factory.NewDbFactorySimple(ctx, dbType)
	dbEng, err := dbFactory.GetDbClient()
	if err != nil {
		return nil
	}
	return &RepoCredentialSimple{
		ctx:      ctx,
		dbEngine: dbEng,
	}
}

func (r *RepoCredentialSimple) AddUser(user repo_model.User) error {

	err := r.dbEngine.Save(db.UserTableName, &user)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoCredentialSimple) GetUser(userId string) (*repo_model.User, error) {
	keySet := db.KeySet{
		db.KeyName(strings.ToLower(string(db.UserIdFieldName))): userId,
	}
	var repoUser repo_model.User
	user, err := r.dbEngine.Get(db.UserTableName, keySet, &repoUser)
	if err != nil {
		return nil, err
	}
	switch user.(type) {
	case *repo_model.User:
		return user.(*repo_model.User), nil
	default:
		return nil, errors.New("unknown data type")
	}

}
func (r *RepoCredentialSimple) AddCredential(credential repo_model.Credential) error {
	err := r.dbEngine.Save(db.CredentialTableName, &credential)
	if err != nil {
		return err
	}
	return nil
}
