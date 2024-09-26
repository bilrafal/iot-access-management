package core

type DoorId string

type WhiteList struct {
	Id           string
	DoorId       DoorId
	CredentialId CredentialId
}

func NewWhiteList(id string, doorId DoorId, credentialId CredentialId) *WhiteList {
	return &WhiteList{Id: id, DoorId: doorId, CredentialId: credentialId}
}

type AccessRequest struct {
	DoorId     DoorId
	Credential CredentialVal
}

func NewAccessRequest(doorId DoorId, credential CredentialVal) *AccessRequest {
	return &AccessRequest{DoorId: doorId, Credential: credential}
}
