package repo

type WhiteList struct {
	Id           string
	DoorId       string
	CredentialId string
}

func NewWhiteList(id string, doorId string, credentialId string) *WhiteList {
	return &WhiteList{Id: id, DoorId: doorId, CredentialId: credentialId}
}
