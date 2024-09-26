package core

type UserCredential struct {
	Id           string
	UserId       UserId
	CredentialId CredentialId
}

func NewUserCredential(id string, userId UserId, credentialId CredentialId) *UserCredential {
	return &UserCredential{Id: id, UserId: userId, CredentialId: credentialId}
}
