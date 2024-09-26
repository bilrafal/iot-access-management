package repo_credential_simple

import (
	"context"
	"errors"
	"fmt"
	"iot-access-management/internal/db"
	db_factory "iot-access-management/internal/db/db_factory"
	"iot-access-management/internal/error/trace_error"
	repo_model "iot-access-management/internal/models/repo"
	"iot-access-management/internal/repo"
	"strings"
)

var (
	errUnknownDataType = errors.New("unknown data type")
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

func (r *RepoCredentialSimple) AddUser(user repo_model.User) *trace_error.TraceError {

	err := r.dbEngine.Save(db.UserTableName, &user)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoCredentialSimple) GetUser(userId string) (*repo_model.User, *trace_error.TraceError) {
	keySet := db.KeySet{
		db.KeyName(strings.ToLower(string(db.UserFkIdFieldName))): userId,
	}
	user, err := r.dbEngine.Get(db.UserTableName, keySet)
	if err != nil {
		return nil, err
	}
	switch user.(type) {
	case *repo_model.User:
		return user.(*repo_model.User), nil
	default:
		return nil, trace_error.ErrUnexpected.From(errUnknownDataType)
	}

}

func (r *RepoCredentialSimple) AddCredential(credential repo_model.Credential) *trace_error.TraceError {
	err := r.dbEngine.Save(db.CredentialTableName, &credential)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoCredentialSimple) GetCredential(credentialId string) (*repo_model.Credential, *trace_error.TraceError) {
	keySet := db.KeySet{
		db.KeyName(strings.ToLower(string(db.IdFieldName))): credentialId,
	}
	credential, err := r.dbEngine.Get(db.CredentialTableName, keySet)
	if err != nil {
		return nil, err
	}
	switch credential.(type) {
	case *repo_model.Credential:
		return credential.(*repo_model.Credential), nil
	default:
		return nil, trace_error.ErrUnexpected.From(errUnknownDataType)
	}

}

func (r *RepoCredentialSimple) ListCredentials() ([]repo_model.Credential, *trace_error.TraceError) {

	listResult, dbErr := r.dbEngine.List(db.CredentialTableName)
	if dbErr != nil {
		return nil, dbErr
	}
	arr, ok := listResult.([]any)
	if !ok {
		return nil, trace_error.ErrUnexpected.From(fmt.Errorf("unexpected type: %T", listResult))
	}
	var result []repo_model.Credential
	for _, wl := range arr {
		d, ok := wl.(*repo_model.Credential)
		if !ok {
			return nil, trace_error.ErrUnexpected.From(fmt.Errorf("unexpected type: %T", listResult))
		}
		result = append(result, *d)
	}
	return result, nil

}

func (r *RepoCredentialSimple) AddUserCredential(userCredential repo_model.UserCredential) *trace_error.TraceError {
	err := r.dbEngine.Save(db.UserCredentialRelTableName, &userCredential)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepoCredentialSimple) GetUserCredentials(userId string) ([]*repo_model.UserCredential, *trace_error.TraceError) {
	keySet := db.KeySet{
		db.KeyName(strings.ToLower(string(db.UserFkIdFieldName))): userId,
	}
	userCredential, err := r.dbEngine.Get(db.UserCredentialRelTableName, keySet)
	if err != nil {
		return nil, err
	}
	switch userCredential.(type) {
	case []*repo_model.UserCredential:
		return userCredential.([]*repo_model.UserCredential), nil
	default:
		return nil, trace_error.ErrUnexpected.From(errUnknownDataType)
	}

}
