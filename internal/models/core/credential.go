package core

import "github.com/google/uuid"

type CredentialId string
type CredentialVal string

type Credential struct {
	Id         CredentialId
	Credential CredentialVal
}

func NewCredential(credential string) *Credential {
	id := uuid.New().String()
	return &Credential{Id: CredentialId(id),Credential: CredentialVal(credential)}
}
