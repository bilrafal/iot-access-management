package repo

type UserCredential struct {
	Id           string //userId
	CredentialId string
}

func NewUserCredential(userId string, credentialId string) *UserCredential {
	return &UserCredential{Id: userId, CredentialId: credentialId}
}
