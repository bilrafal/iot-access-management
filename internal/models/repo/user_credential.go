package repo

type UserCredential struct {
	UserId       string
	CredentialId string
}

func NewUserCredential(userId string, credentialId string) *UserCredential {
	return &UserCredential{UserId: userId, CredentialId: credentialId}
}
