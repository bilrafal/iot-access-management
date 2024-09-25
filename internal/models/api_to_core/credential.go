package api_to_core

import (
	"iot-access-management/internal/models/api"
	"iot-access-management/internal/models/core"
)

func ApiCreateCredentialToCoreCredential(apiCredential api.CredentialCreateRequest) core.Credential {
	return core.Credential{
		Credential: core.CredentialVal(apiCredential.Credential),
	}
}

func ApiCredentialResponseToCoreCredential(apiCredential api.CredentialResponse) core.Credential {
	return core.Credential{
		Id:         core.CredentialId(apiCredential.Id),
		Credential: core.CredentialVal(apiCredential.Credential),
	}
}
