package core

type UserCredential struct {
	UserId       UserId
	CredentialId CredentialId
}

func NewUserCredential(userId UserId, credentialId CredentialId) *UserCredential {
	return &UserCredential{UserId: userId, CredentialId: credentialId}
}
