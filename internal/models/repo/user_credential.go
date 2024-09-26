package repo

type UserCredential struct {
	Id           string
	UserId       string
	CredentialId string
}

func NewUserCredential(id string, userId string, credentialId string) *UserCredential {
	return &UserCredential{Id: id, UserId: userId, CredentialId: credentialId}
}
