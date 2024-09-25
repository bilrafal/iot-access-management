package api

type CredentialCreateRequest struct {
	Credential string `json:"code" binding:"required"`
}

func NewCredentialCreateRequest(credential string) *CredentialCreateRequest {
	return &CredentialCreateRequest{Credential: credential}
}

type CredentialResponse struct {
	Id         string `json:"id"`
	Credential string `json:"code"`
}

func NewCredentialResponse(id, credential string) *CredentialResponse {
	return &CredentialResponse{Id: id, Credential: credential}
}
