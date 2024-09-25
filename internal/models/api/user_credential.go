package api

type AssignCredentialToUserRequest struct {
	UserId       string `json:"user_id" binding:"required"`
	CredentialId string `json:"credential_id" binding:"required"`
}

func NewAssignCredentialToUserRequest(userId, credentialId string) *AssignCredentialToUserRequest {
	return &AssignCredentialToUserRequest{UserId: userId, CredentialId: credentialId}
}
