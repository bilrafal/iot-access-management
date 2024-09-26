package repo

type Credential struct {
	CredentialId string
	Code         string
}

func NewCredential(id string, code string) *Credential {
	return &Credential{CredentialId: id, Code: code}
}
