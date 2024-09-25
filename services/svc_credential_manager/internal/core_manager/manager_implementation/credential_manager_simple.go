package manager_implementation

import (
	"context"
	"errors"
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/core_to_repo"
	"iot-access-management/internal/models/repo"
	"iot-access-management/internal/models/repo_to_core"
	"iot-access-management/internal/repo"
	"iot-access-management/internal/util"
	"iot-access-management/services/svc_credential_manager/internal/core_manager"
)

type CredentialManagerSimple struct {
	ctx  context.Context
	repo repo.RepoCredential
}

func NewCredentialManagerSimple(repo repo.RepoCredential) core_manager.CredentialManager {
	return &CredentialManagerSimple{repo: repo}
}

func (cm CredentialManagerSimple) CreateUser(user core.User) (*core.User, error) {
	err := validateUser(user)
	if err != nil {
		return nil, err
	}

	coreUser := core.NewUser(user.Name)
	repoUser := core_to_repo.CoreUserToRepoUser(*coreUser)
	err = cm.repo.AddUser(repoUser)
	if err != nil {
		return nil, err
	}
	return coreUser, nil
}

func validateUser(user core.User) error {
	if util.IsVoidString(user.Name) {
		return errors.New("username is empty")
	}
	return nil
}

func (cm CredentialManagerSimple) GetUser(id core.UserId) (*core.User, error) {
	repoUser, err := cm.repo.GetUser(repo.UserId(id))
	if err != nil {
		return nil, err
	}
	coreUser := repo_to_core.RepoUserToCoreUser(*repoUser)
	return &coreUser, nil
}

func (CredentialManagerSimple) CreateCredential(accessCode string) core.CredentialId {
	//TODO implement me
	panic("implement me")
}

func (CredentialManagerSimple) AssignCredentialToUser(userId string, credId core.CredentialId) error {
	//TODO implement me
	panic("implement me")
}

func (CredentialManagerSimple) AuthorizeUserOnDoor(credId core.CredentialId) error {
	//TODO implement me
	panic("implement me")
}

func (CredentialManagerSimple) RevokeAuthorization(credId core.CredentialId) error {
	//TODO implement me
	panic("implement me")
}
