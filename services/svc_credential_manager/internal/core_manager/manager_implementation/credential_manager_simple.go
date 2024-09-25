package manager_implementation

import (
	"context"
	"errors"
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/core_to_repo"
	repo_model "iot-access-management/internal/models/repo"
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
	//TODO: possibly check if there is no user with the same name
	if util.IsVoidString(user.Name) {
		return errors.New("username is empty")
	}
	return nil
}

func (cm CredentialManagerSimple) GetUser(id core.UserId) (*core.User, error) {
	repoUser, err := cm.repo.GetUser(string(id))
	if err != nil {
		return nil, err
	}
	coreUser := repo_to_core.RepoUserToCoreUser(*repoUser)
	return &coreUser, nil
}

func (cm CredentialManagerSimple) CreateCredential(accessCode string) (core.CredentialId, error) {
	err := validateCredential(accessCode)
	if err != nil {
		return core.VoidCredentialId, err
	}

	coreCredential := core.NewCredential(accessCode)
	repoCredential := core_to_repo.CoreCredentialToRepoCredential(*coreCredential)
	err = cm.repo.AddCredential(repoCredential)
	if err != nil {
		return core.VoidCredentialId, err
	}
	return coreCredential.Id, nil

}

func validateCredential(accessCode string) error {
	//TODO: refactor magic number
	if len(accessCode) != 8 {
		return errors.New("accessCode length invalid")
	}
	return nil
}

func (cm CredentialManagerSimple) GetCredential(id core.CredentialId) (*core.Credential, error) {
	repoCredential, err := cm.repo.GetCredential(string(id))
	if err != nil {
		return nil, err
	}
	coreCredential := repo_to_core.RepoCredentialToCoreCredential(*repoCredential)
	return &coreCredential, nil
}
func (cm CredentialManagerSimple) AssignCredentialToUser(userId core.UserId, credId core.CredentialId) error {

	user, _ := cm.GetUser(userId)
	if user == nil {
		return core_manager.ErrUserNotFound.Generate()
	}

	credential, _ := cm.GetCredential(credId)
	if credential == nil {
		return core_manager.ErrCredentialNotFound.Generate()
	}

	userCredential := repo_model.NewUserCredential(string(userId), string(credId))
	err := cm.repo.AddUserCredential(*userCredential)
	if err != nil {
		return err
	}

	return nil
}

func (cm CredentialManagerSimple) GetUserCredentials(userId core.UserId) ([]*core.UserCredential, error) {
	repoCredentials, err := cm.repo.GetUserCredentials(string(userId))
	if err != nil {
		return nil, err
	}

	var coreCredentials []*core.UserCredential
	for _, credential := range repoCredentials {
		coreCredential := repo_to_core.RepoUserCredentialToCoreUserCredential(*credential)
		coreCredentials = append(coreCredentials, &coreCredential)
	}

	return coreCredentials, nil
}

func (CredentialManagerSimple) AuthorizeUserOnDoor(credId core.CredentialId) error {
	//TODO implement me
	panic("implement me")
}

func (CredentialManagerSimple) RevokeAuthorization(credId core.CredentialId) error {
	//TODO implement me
	panic("implement me")
}
