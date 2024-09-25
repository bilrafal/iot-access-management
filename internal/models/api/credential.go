package api

type CredentialCreateRequest struct {
	Credential string `json:"credential" binding:"required"`
}

func NewCredentialCreateRequest(credential string) *CredentialCreateRequest {
	return &CredentialCreateRequest{Credential: credential}
}

type CredentialCreateResponse struct {
	Id string `json:"id" binding:"required"`
}

func NewCredentialCreateResponse(credentialId string) *CredentialCreateResponse {
	return &CredentialCreateResponse{Id: credentialId}
}

type CredentialResponse struct {
	Id         string `json:"id"`
	Credential string `json:"code"`
}

func NewCredentialResponse(id, credential string) *CredentialResponse {
	return &CredentialResponse{Id: id, Credential: credential}
}
