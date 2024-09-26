package manager_implementation

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"iot-access-management/internal/client"
	"iot-access-management/internal/error/trace_error"
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/core_to_repo"
	repo_model "iot-access-management/internal/models/repo"
	"iot-access-management/internal/models/repo_to_core"
	"iot-access-management/internal/repo"
	"iot-access-management/internal/util"
	"iot-access-management/services/svc_credential_manager/internal/core_manager"
	"net/http"
)

type CredentialManagerSimple struct {
	ctx       context.Context
	repo      repo.RepoCredential
	iotClient client.Client
}

func NewCredentialManagerSimple(repo repo.RepoCredential, iotClient client.Client) core_manager.CredentialManager {
	return &CredentialManagerSimple{repo: repo, iotClient: iotClient}
}

func (cm CredentialManagerSimple) CreateUser(user core.User) (*core.User, *trace_error.TraceError) {
	err := validateUser(user)
	if err != nil {
		return nil, trace_error.ErrUnexpected.From(err)
	}

	coreUser := core.NewUser(user.Name)
	repoUser := core_to_repo.CoreUserToRepoUser(*coreUser)
	repoErr := cm.repo.AddUser(repoUser)
	if repoErr != nil {
		return nil, repoErr
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

func (cm CredentialManagerSimple) GetUser(id core.UserId) (*core.User, *trace_error.TraceError) {
	repoUser, err := cm.repo.GetUser(string(id))
	if err != nil {
		return nil, err
	}
	coreUser := repo_to_core.RepoUserToCoreUser(*repoUser)
	return &coreUser, nil
}

func (cm CredentialManagerSimple) CreateCredential(accessCode string) (core.CredentialId, *trace_error.TraceError) {
	err := validateCredential(accessCode)
	if err != nil {
		return core.VoidCredentialId, trace_error.ErrUnexpected.From(err)
	}

	coreCredential := core.NewCredential(accessCode)
	repoCredential := core_to_repo.CoreCredentialToRepoCredential(*coreCredential)
	repoErr := cm.repo.AddCredential(repoCredential)
	if repoErr != nil {
		return core.VoidCredentialId, repoErr
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

func (cm CredentialManagerSimple) GetCredential(id core.CredentialId) (*core.Credential, *trace_error.TraceError) {

	repoCredential, err := cm.repo.GetCredential(string(id))
	if err != nil {
		return nil, err
	}
	coreCredential := repo_to_core.RepoCredentialToCoreCredential(*repoCredential)
	return &coreCredential, nil
}

func (cm CredentialManagerSimple) GetCredentialIdByCode(credentialCode core.CredentialVal) (
	*core.Credential, *trace_error.TraceError,
) {
	validatioErr := validateCredential(string(credentialCode))
	if validatioErr != nil {
		return nil, core_manager.ErrUnexpected.From(validatioErr)
	}
	repoCredential, err := cm.repo.ListCredentials()
	if err != nil {
		return nil, err
	}

	var coreCredentialPtr *core.Credential
	for _, credential := range repoCredential {
		if credential.Code == string(credentialCode) {
			coreCredential := repo_to_core.RepoCredentialToCoreCredential(credential)
			coreCredentialPtr = &coreCredential
			break
		}
	}

	if coreCredentialPtr == nil {
		return nil, core_manager.ErrCredentialNotFound.Generate()
	}
	return coreCredentialPtr, nil
}

func (cm CredentialManagerSimple) AssignCredentialToUser(userId core.UserId, credId core.CredentialId) *trace_error.TraceError {

	user, _ := cm.GetUser(userId)
	if user == nil {
		return core_manager.ErrUserNotFound.Generate()
	}

	credential, _ := cm.GetCredential(credId)
	if credential == nil {
		return core_manager.ErrCredentialNotFound.Generate()
	}

	id := uuid.NewString()
	userCredential := repo_model.NewUserCredential(id, string(userId), string(credId))
	err := cm.repo.AddUserCredential(*userCredential)
	if err != nil {
		return err
	}

	return nil
}

func (cm CredentialManagerSimple) GetUserCredentials(userId core.UserId) ([]*core.UserCredential, *trace_error.TraceError) {
	repoCredentials, repoErr := cm.repo.GetUserCredentials(string(userId))
	if repoErr != nil {
		return nil, repoErr
	}

	var coreCredentials []*core.UserCredential
	for _, credential := range repoCredentials {
		coreCredential := repo_to_core.RepoUserCredentialToCoreUserCredential(*credential)
		coreCredentials = append(coreCredentials, &coreCredential)
	}

	return coreCredentials, nil
}

func (cm CredentialManagerSimple) AuthorizeUserOnDoor(doorId core.DoorId, credId core.CredentialId) *trace_error.TraceError {
	req := api.NewWhiteListCreateRequest(string(doorId), string(credId))
	reqAsBytes, err := json.Marshal(req)
	if err != nil {
		return trace_error.ErrUnexpected.From(err)
	}
	body := bytes.NewReader(reqAsBytes)
	httpReq, reqErr := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/white-list", cm.iotClient.Host, cm.iotClient.Port), body)
	if reqErr != nil {
		return trace_error.ErrUnexpected.From(reqErr)
	}
	httpClient := &http.Client{}
	resp, respErr := httpClient.Do(httpReq)
	if respErr != nil {
		return trace_error.ErrUnexpected.From(respErr)
	}

	defer resp.Body.Close()
	fmt.Println("resp status", resp.Status)

	data, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return trace_error.ErrUnexpected.From(readErr)
	}
	fmt.Println("resp body", string(data))

	return nil
}

func (cm CredentialManagerSimple) RevokeAuthorization(doorId core.DoorId, credId core.CredentialId) *trace_error.TraceError {
	req := api.NewWhiteListCreateRequest(string(doorId), string(credId))
	reqAsBytes, err := json.Marshal(req)
	if err != nil {
		return trace_error.ErrUnexpected.From(err)
	}
	body := bytes.NewReader(reqAsBytes)

	httpReq, reqErr := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s:%d/white-list", cm.iotClient.Host, cm.iotClient.Port), body)
	if reqErr != nil {
		return trace_error.ErrUnexpected.From(reqErr)
	}

	httpClient := &http.Client{}
	resp, respErr := httpClient.Do(httpReq)
	if respErr != nil {
		return trace_error.ErrUnexpected.From(respErr)
	}

	defer resp.Body.Close()
	fmt.Println("resp status", resp.Status)

	_, readBodyErr := io.ReadAll(resp.Body)
	if readBodyErr != nil {
		return trace_error.ErrUnexpected.From(readBodyErr)
	}

	return nil
}
