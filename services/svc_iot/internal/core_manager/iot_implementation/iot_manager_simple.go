package iot_implementation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"iot-access-management/internal/client"
	"iot-access-management/internal/error/trace_error"
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
	"iot-access-management/internal/models/core_to_repo"
	"iot-access-management/internal/models/repo_to_core"
	"iot-access-management/internal/repo"
	"iot-access-management/services/svc_iot/internal/core_manager"
	"net/http"
)

type IoTManagerSimple struct {
	ctx      context.Context
	repo     repo.IotRepo
	cmClient client.Client
}

func NewCredentialManagerSimple(repo repo.IotRepo, cmClient client.Client) core_manager.IoTManager {
	return &IoTManagerSimple{repo: repo, cmClient: cmClient}
}

func (iotms *IoTManagerSimple) CreateWhiteList(whitelist core.WhiteList) *trace_error.TraceError {

	whitelist.Id = uuid.NewString()
	repoWhiteList := core_to_repo.CoreWhiteListToRepoWhiteList(whitelist)
	err := iotms.repo.CreateWhiteList(repoWhiteList)
	if err != nil {
		return err
	}
	return nil
}

func (iotms *IoTManagerSimple) DeleteWhiteList(whitelist core.WhiteList) *trace_error.TraceError {

	items, listErr := iotms.GetWhiteList()
	if listErr != nil {
		return core_manager.ErrWhiteListNotFound.From(
			fmt.Errorf("no element found for doorId [%s] and credentialId [%s]", whitelist.DoorId, whitelist.CredentialId))
	}

	var found *core.WhiteList
	for _, item := range items {
		if item.DoorId == whitelist.DoorId && item.CredentialId == whitelist.CredentialId {
			found = &item
			break
		}
	}
	if found == nil {
		return core_manager.ErrWhiteListNotFound.From(
			fmt.Errorf("no element found for door_Id [%s] and credential_Id [%s]", whitelist.DoorId, whitelist.CredentialId))
	}

	repoWhiteList := core_to_repo.CoreWhiteListToRepoWhiteList(*found)
	err := iotms.repo.DeleteWhiteList(repoWhiteList)
	if err != nil {
		return err
	}
	return nil
}

func (iotms *IoTManagerSimple) GetWhiteList() ([]core.WhiteList, *trace_error.TraceError) {

	repoWhiteList, err := iotms.repo.ListWhiteList()
	if err != nil {
		return nil, core_manager.ErrWhiteListNotFound.From(errors.New("No elements found for whitelist"))
	}

	var result []core.WhiteList
	for _, wl := range repoWhiteList {
		result = append(result, repo_to_core.RepoWhiteListToCoreWhiteList(wl))
	}

	return result, nil
}

func (iotms *IoTManagerSimple) RequestAccess(accessRequest core.AccessRequest) (bool, *trace_error.TraceError) {
	credential, err := iotms.getCredentialIdByCredential(accessRequest.Credential)
	if err != nil {
		return false, err
	}

	allWhiteLists, trErr := iotms.GetWhiteList()
	if trErr != nil {
		return false, trErr
	}

	found := false
	for _, wl := range allWhiteLists {
		if wl.DoorId == accessRequest.DoorId && wl.CredentialId == credential.Id {
			found = true
			break
		}
	}

	return found, nil
}

func (iotms *IoTManagerSimple) getCredentialIdByCredential(cred core.CredentialVal) (*core.Credential, *trace_error.TraceError) {
	httpReq, reqErr := http.NewRequest(http.MethodGet,
		fmt.Sprintf("http://%s:%d/credential/%s", iotms.cmClient.Host, iotms.cmClient.Port, cred), nil)
	if reqErr != nil {
		return nil, trace_error.ErrUnexpected.From(reqErr)
	}
	httpClient := &http.Client{}
	resp, respErr := httpClient.Do(httpReq)
	if respErr != nil {
		return nil, trace_error.ErrUnexpected.From(respErr)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, trace_error.ErrNotFound.Generate()
	}
	var apiCredential api.CredentialCreateResponse
	err := json.NewDecoder(resp.Body).Decode(&apiCredential)
	if err != nil {
		return nil, trace_error.ErrUnexpected.From(err)
	}

	return &core.Credential{Id: core.CredentialId(apiCredential.Id), Credential: cred}, nil
}
