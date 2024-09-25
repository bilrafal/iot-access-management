package repo_credential_simple

import (
	"context"
	"errors"
	"iot-access-management/internal/db"
	db_factory "iot-access-management/internal/db/db_factory"
	"iot-access-management/internal/models/repo"
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

func (r *RepoCredentialSimple) AddUser(user repo.User) error {

	err := r.dbEngine.Save(db.UserTableName, &user)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoCredentialSimple) GetUser(userId repo.UserId) (*repo.User, error) {
	keySet := db.KeySet{
		db.KeyName(strings.ToLower(string(db.UserIdFieldName))): userId,
	}
	var repoUser repo.User
	user, err := r.dbEngine.Get(db.UserTableName, keySet, &repoUser)
	if err != nil {
		return nil, err
	}
	switch user.(type) {
	case *repo.User:
		return user.(*repo.User), nil
	default:
		return nil, errors.New("unknown data type")
	}

}
func (r *RepoCredentialSimple) AddCredential(accessCode string) error {
	return nil
}
