package api

type WhiteListCreateRequest struct {
	DoorId       string `json:"door_id" binding:"required"`
	CredentialId string `json:"credential_id" binding:"required"`
}

func NewWhiteListCreateRequest(doorId, credentialId string) *WhiteListCreateRequest {
	return &WhiteListCreateRequest{DoorId: doorId, CredentialId: credentialId}
}

type AccessRequest struct {
	DoorId     string `json:"door_id" binding:"required"`
	Credential string `json:"credential" binding:"required"`
}

func NewAccessRequest(doorId string, credential string) *AccessRequest {
	return &AccessRequest{DoorId: doorId, Credential: credential}
}
